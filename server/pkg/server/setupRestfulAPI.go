package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/klopjq/telemedicine/config"
	"github.com/klopjq/telemedicine/internal/log"
)

type SetupRestfulAPI struct {
	*config.Config
	Logger  log.Logger
	Ctx     context.Context
	DB      *sqlx.DB
	Handler http.Handler
}

func (s SetupRestfulAPI) GetAddr() string {
	//return fmt.Sprintf(":%d", s.Config.ServerPort)
	return ":" + strconv.Itoa(s.Config.ServerPort)
}

func (s SetupRestfulAPI) GetLogger() log.Logger {
	return s.Logger
}

func (s SetupRestfulAPI) GetHttpHandler() http.Handler {
	return s.Handler
}

func (s SetupRestfulAPI) GetDatabase() *sqlx.DB {
	return s.DB
}

func (s SetupRestfulAPI) GetConfig() *config.Config {
	return s.Config
}

func (s SetupRestfulAPI) GetContext() context.Context {
	return s.Ctx
}
