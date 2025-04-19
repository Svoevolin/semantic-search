package config

import (
	"fmt"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type App struct {
	DB
	Logger
}

type DB struct {
	Adapter            string `env:"DB_ADAPTER, notEmpty"`
	Host               string `env:"DB_HOST, notEmpty"`
	Port               string `env:"DB_PORT, notEmpty"`
	User               string `env:"DB_USER, notEmpty"`
	Password           string `env:"DB_PASSWORD, notEmpty"`
	Name               string `env:"DB_NAME, notEmpty"`
	SSLMode            string `env:"DB_SSL_MODE, notEmpty"`
	Timezone           string `env:"DB_TIMEZONE, notEmpty"`
	LogLevel           string `env:"LOG_LEVEL, notEmpty"`
	TargetSessionAttrs string `env:"DB_TARGET_SESSION_ATTRS, notEmpty"`
}

type Logger struct {
	Level   string `env:"LOG_LEVEL, notEmpty"`
	Format  string `env:"LOG_FORMAT, notEmpty"`
	Pretty  bool   `env:"LOG_PRETTY_ENABLE, notEmpty"`
	Discard bool   `env:"LOG_DISCARD_ENABLE, notEmpty"`
}

func New[T any](files ...string) (T, error) {
	var cfg T

	for idx, file := range files {
		abs, err := filepath.Abs(file)
		if err != nil {
			return cfg, fmt.Errorf("failed to resolve absolute path for evv file %s: %w", file, err)
		}
		files[idx] = abs
	}

	_ = godotenv.Load(files...)
	return cfg, env.Parse(&cfg)
}
