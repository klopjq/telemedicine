package main

import (
	"context"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/klopjq/telemedicine/config"
	"github.com/klopjq/telemedicine/internal/data"
	"github.com/klopjq/telemedicine/internal/log"
	"github.com/klopjq/telemedicine/server/pkg/routing"
	"github.com/klopjq/telemedicine/server/pkg/server"
)

type serverSetup struct {
	*config.Config
	logger  log.Logger
	ctx     context.Context
	db      *sqlx.DB
	handler http.Handler
}

func (s serverSetup) GetAddr() string {
	//return fmt.Sprintf(":%d", s.Config.ServerPort)
	return ":" + strconv.Itoa(s.Config.ServerPort)
}

func (s serverSetup) GetLogger() log.Logger {
	return s.logger
}

func (s serverSetup) GetHttpHandler() http.Handler {
	return s.handler
}

func (s serverSetup) GetDatabase() *sqlx.DB {
	return s.db
}

func (s serverSetup) GetConfig() *config.Config {
	return s.Config
}

func (s serverSetup) GetContext() context.Context {
	return s.ctx
}

const (
	version = 1.0
)

func main() {
	numProcessor := runtime.NumCPU()
	_ = os.Setenv("CONFIG_PATH", "../../../config/local.yaml")
	configPath := os.Getenv("CONFIG_PATH")

	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	logger := log.New().With(ctx, "version", version)
	cfg, err := config.New(configPath, logger)
	if err != nil {
		panic(err)
	}

	db := data.DB{}
	if err := db.Open(ctx, "postgres", cfg.LocalPgSqlDsn, 50); err != nil {
		panic(err)
	}
	defer db.Close()
	setup := serverSetup{
		Config:  cfg,
		logger:  logger,
		ctx:     ctx,
		db:      db.Connection,
		handler: routing.New(logger),
	}
	restServer, err := server.NewRestfulAPI(setup)

	if err != nil {
		panic(err)
	}

	logger.Infof("Processors: %d", numProcessor)
	if err := restServer.Run(); err != nil {
		panic(err)
	}
}
