package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/klopjq/telemedicine/config"
	"github.com/klopjq/telemedicine/internal/log"
)

type Setup interface {
	GetAddr() string
	GetLogger() log.Logger
	GetHttpHandler() http.Handler
	GetDatabase() *sqlx.DB
	GetConfig() *config.Config
	GetContext() context.Context
}

type Server struct {
	httpServer *http.Server
	logger     log.Logger
	ctx        context.Context
	db         *sqlx.DB
	cfg        *config.Config
}

func NewRestfulAPI(setup Setup) (*Server, error) {
	return &Server{
		ctx:    setup.GetContext(),
		logger: setup.GetLogger(),
		cfg:    setup.GetConfig(),
		db:     setup.GetDatabase(),
		httpServer: &http.Server{
			Addr:           setup.GetAddr(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        setup.GetHttpHandler(),
		},
	}, nil
}

func (s *Server) Run() error {
	errc := make(chan error, 1)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			errc <- err
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	select {
	case err := <-errc:
		return err
	case <-quit:
		s.logger.Info("quit signal received")
		ctx, shutdown := context.WithTimeout(s.ctx, 5*time.Second)
		defer shutdown()
		return s.httpServer.Shutdown(ctx)
	}
}
