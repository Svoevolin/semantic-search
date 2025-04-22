package app

import (
	"context"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
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
	c.Config = &cfg
	c.Logger = sl.NewLogger(c.Config)

	// Adapter
	//searchClient := searcher.NewClient(http.DefaultClient, c.Logger, c.Config)

	// Service
	//c.DocumentService = service.NewDocument(searchClient)

	return nil
}
