package config

import (
	"errors"
	"io/ioutil"

	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v3"

	"github.com/klopjq/telemedicine/internal/log"
)

const (
	defaultServerPort = 8000
)

type Config struct {
	ServerPort       int    `yaml:"ServerPort" env:"PORT"`
	ServerApiKey     string `yaml:"ServerApiKey" env:"SERVER_API_KEY"`
	LocalPgSqlDsn    string `yaml:"LocalPgSqlDsn" env:"DATABASE_URL"`
	DbMaxConnections int    `yaml:"DbMaxConnections" env:"DB_MAX_CONNECTIONS"`
	ServerPem        string `yaml:"ServerPem" env:"SERVER_PEM"`
	ServerKey        string `yaml:"ServerKey" env:"SERVER_KEY"`
}

func (c Config) Validate() error {
	if c.ServerPort == 0 {
		return errors.New("provide and non zero SERVER_PORT")
	}
	if c.ServerApiKey == "" {
		return errors.New("provide a API key server SERVER_API_KEY")
	}
	if c.LocalPgSqlDsn == "" {
		return errors.New("provide a database LOCAL_PGSQL_DSN")
	}
	if c.DbMaxConnections == 0 {
		return errors.New("provide and non zero DB_MAX_CONNECTIONS")
	}
	return nil
}

func New(fname string, logger log.Logger) (*Config, error) {
	cfg := Config{
		ServerPort: defaultServerPort,
	}

	if fname != "" {
		b, err := ioutil.ReadFile(fname)
		if err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(b, &cfg); err != nil {
			return nil, err
		}
	}

	if err := env.New("", logger.Infof).Load(&cfg); err != nil {
		return nil, err
	}

	return &cfg, cfg.Validate()
}

func (c *Config) GetDsn(set string) (dsn string) {
	return c.LocalPgSqlDsn
}
