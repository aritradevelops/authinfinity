package config

import (
	"github.com/caarlos0/env/v11"
)

const (
	emailVerificationHashLength = 32
	refreshTokenHashLength      = 64
)

type Env struct {
	Port                        int    `env:"PORT" envDefault:"8080"`
	DBConnectionURI             string `env:"DB_CONNECTION_URI,required"`
	DBKind                      int    `env:"DB_KIND,required"`
	ServiceName                 string `env:"SERVICE_NAME" envDefault:"Authone Backend"`
	BaseDomain                  string `env:"BASE_DOMAIN,required"`
	EmailVerificationHashExpiry string `env:"EMAIL_VERIFICATION_HASH_EXPIRY" envDefault:"15m"`
}

type Config struct {
	Env                         Env
	EmailVerificationHashLength int
	RefreshTokenHashLength      int
}

var (
	instance Config
)

func Load() (Config, error) {
	var config Config
	err := env.Parse(&config.Env)

	config.EmailVerificationHashLength = emailVerificationHashLength
	config.RefreshTokenHashLength = refreshTokenHashLength

	instance = config
	return config, err
}

func Instance() Config {
	return instance
}
