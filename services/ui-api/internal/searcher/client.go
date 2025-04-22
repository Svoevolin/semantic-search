package searcher

import (
	"context"
	"net/http"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type Client struct {
	cfg    *config.App
	client *http.Client
	logger logger.Logger
}

func NewClient(client *http.Client, logger logger.Logger, cfg *config.App) *Client {
	return &Client{
		client: client,
		logger: logger,
		cfg:    cfg,
	}
}

func (c *Client) Search(ctx context.Context, query string) ([]domain.Document, error) {
	const op = "adapter.searcher.client.Search"
	_ = op
	return []domain.Document{}, nil
}
