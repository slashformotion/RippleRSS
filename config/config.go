package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// http server
	Host string `env:"RIPPLERSS_HOST"  envDefault:"0.0.0.0"`
	Port uint16 `env:"RIPPLERSS_PORT"  envDefault:"12000"`

	// RSS option
	RefreshInterval time.Duration `env:"RIPPLERSS_DEFAULT_REFRESH_INTERVAL"  envDefault:"10s"`

	// DB
	DbPath string `env:"RIPPLERSS_DB_PATH"  envDefault:"./db.db"`
}

func LoadFromEnvVariables() (Config, error) {
	return env.ParseAs[Config]()
}
