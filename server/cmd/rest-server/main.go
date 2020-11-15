package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/klopjq/telemedicine/config"
	"github.com/klopjq/telemedicine/internal/data"
	"github.com/klopjq/telemedicine/internal/log"
	"github.com/klopjq/telemedicine/server/pkg/routing"
	"github.com/klopjq/telemedicine/server/pkg/server"
)

const (
	version = 1.0
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	numProcessor := runtime.NumCPU()
	_ = os.Setenv("CONFIG_PATH", "../../../config/local.yaml")
	configPath := os.Getenv("CONFIG_PATH")

	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		ctxCancel()
	}()

	logger := log.New().With(ctx, "version", version)
	cfg, err := config.New(configPath, logger)
	if err != nil {
		return err
	}

	db := data.DB{}
	if err := db.Open(ctx, "postgres", cfg.LocalPgSqlDsn, 50); err != nil {
		return err
	}
	defer db.Close()

	restServer, err := server.NewRestfulAPI(server.SetupRestfulAPI{
		Config:  cfg,
		Logger:  logger,
		Ctx:     ctx,
		DB:      db.Connection,
		Handler: routing.New(logger),
	})

	if err != nil {
		return err
	}

	logger.Infof("Processors: %d", numProcessor)
	if err := restServer.Run(); err != nil {
		return err
	}
	return nil
}
