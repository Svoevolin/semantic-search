package httpclient

import (
	"net/http"
	"time"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type RoundTripperDecorator func(http.RoundTripper) http.RoundTripper

type BuilderConfig struct {
	Logging bool
}

type Builder struct {
	transportDecorators []RoundTripperDecorator
	config              BuilderConfig
}

func NewBuilder(config BuilderConfig) *Builder {
	return &Builder{
		config: config,
	}
}

func (f *Builder) WithLogging(logger logger.Logger) *Builder {
	if f.config.Logging {
		f.transportDecorators = append(f.transportDecorators, LoggingWrapper(logger))
	}
	return f
}

func (f *Builder) WithRequestID() *Builder {
	f.transportDecorators = append(f.transportDecorators, RequestIDWrapper())
	return f
}

type Config struct {
	Transport     http.RoundTripper
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	Timeout       time.Duration
}

func (f *Builder) Create(cfg Config) *http.Client {
	transport := cfg.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	for _, decorator := range f.transportDecorators {
		transport = decorator(transport)
	}
	return &http.Client{
		Transport:     transport,
		CheckRedirect: cfg.CheckRedirect,
		Jar:           cfg.Jar,
		Timeout:       cfg.Timeout,
	}
}
