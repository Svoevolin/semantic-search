package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type App struct {
	LoggingOutgoingReqEnable bool `env:"UI_API_LOGGING_OUTGOING_REQUESTS_ENABLE" envDefault:"true"`

	Logger
	PublicServer
	Minio
	Searcher
	Kafka
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

type Searcher struct {
	Timeout time.Duration `env:"UI_API_SEARCHER_HTTPCLIENT_TIMEOUT" envDefault:"10s"`
}

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

type Minio struct {
	Endpoint  string `env:"UI_API_MINIO_ENDPOINT,notEmpty"`
	AccessKey string `env:"UI_API_MINIO_ACCESS_KEY,notEmpty"`
	SecretKey string `env:"UI_API_MINIO_SECRET_KEY,notEmpty"`
	Bucket    string `env:"UI_API_MINIO_BUCKET,notEmpty"`
	UseSSL    bool   `env:"UI_API_MINIO_USE_SSL,notEmpty"`
}

type Kafka struct {
	Broker string `env:"UI_API_KAFKA_BROKER,notEmpty"`
	Topic  string `env:"UI_API_KAFKA_TOPIC,notEmpty"`
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
