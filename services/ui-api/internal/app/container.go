package app

import (
	"context"
	"fmt"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/adapter"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/adapter/httpclient"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/service"
)

// Container is alphabetically ordered
type Container struct {
	Config          *config.App
	Logger          logger.Logger
	DocumentService domain.DocumentService
}

func NewContainer(ctx context.Context, cfg config.App) (*Container, error) {
	c := &Container{}
	err := c.initContainer(ctx, cfg)
	return c, err
}

//nolint:unparam
func (c *Container) initContainer(_ context.Context, cfg config.App) error {
	const op = "internal.app.container.initContainer"

	c.Config = &cfg
	c.Logger = sl.NewLogger(c.Config)

	clientBuilder := httpclient.NewBuilder(httpclient.BuilderConfig{Logging: c.Config.LoggingOutgoingReqEnable}).
		WithLogging(c.Logger).
		WithRequestID()

	// Adapter
	searchClient := adapter.NewSearcherClient(clientBuilder.Create(
		httpclient.Config{Timeout: c.Config.Searcher.Timeout}), c.Logger, c.Config,
	)

	storageClient, err := adapter.NewMinioClient(c.Logger, c.Config)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Service
	c.DocumentService = service.NewDocument(searchClient, storageClient, c.Logger)

	return nil
}
