package config

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/caarlos0/env/v11"
)

type Env struct {
	Port            int             `env:"PORT" envDefault:"8080"`
	DBConnectionURI string          `env:"DB_CONNECTION_URI,required"`
	DBKind          db.DatabaseKind `env:"DB_KIND,required"`
	ServiceName     string          `env:"SERVICE_NAME" envDefault:"Authone Backend"`
}

type Config struct {
	Env Env
}

func Load() (Config, error) {
	var config Config
	err := env.Parse(&config.Env)
	return config, err
}
