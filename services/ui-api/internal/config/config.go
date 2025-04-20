package config

import (
	"fmt"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type App struct {
	Logger
	PublicServer
}

//type DB struct {
//	Adapter            string `env:"DB_ADAPTER, notEmpty"`
//	Host               string `env:"DB_HOST, notEmpty"`
//	Port               string `env:"DB_PORT, notEmpty"`
//	User               string `env:"DB_USER, notEmpty"`
//	Password           string `env:"DB_PASSWORD, notEmpty"`
//	Name               string `env:"DB_NAME, notEmpty"`
//	SSLMode            string `env:"DB_SSL_MODE, notEmpty"`
//	Timezone           string `env:"DB_TIMEZONE, notEmpty"`
//	LogLevel           string `env:"LOG_LEVEL, notEmpty"`
//	TargetSessionAttrs string `env:"DB_TARGET_SESSION_ATTRS, notEmpty"`
//}

type Logger struct {
	Level   string `env:"UI_API_LOG_LEVEL,notEmpty"`
	Format  string `env:"UI_API_LOG_FORMAT,notEmpty"`
	Pretty  bool   `env:"UI_API_LOG_PRETTY_ENABLE,notEmpty"`
	Discard bool   `env:"UI_API_LOG_DISCARD_ENABLE,notEmpty"`
}

type PublicServer struct {
	Port string `env:"UI_API_PUBLIC_SERVER_PORT,notEmpty"`
	CORSConfig
}

type CORSConfig struct {
	AllowOrigins []string `env:"UI_API_PUBLIC_SERVER_CORS_ALLOW_ORIGINS,notEmpty" envSeparator:","`
	AllowMethods []string `env:"UI_API_PUBLIC_SERVER_CORS_ALLOW_METHODS,notEmpty" envSeparator:","`
}

func New[T any](files ...string) (T, error) {
	var cfg T

	for idx, file := range files {
		abs, err := filepath.Abs(file)
		if err != nil {
			return cfg, fmt.Errorf("failed to resolve absolute path for env file %s: %w", file, err)
		}
		files[idx] = abs
	}

	_ = godotenv.Overload(files...)
	return cfg, env.Parse(&cfg)
}
