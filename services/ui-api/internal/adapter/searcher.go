package adapter

import (
	"context"
	"net/http"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

// compile-time check
var _ domain.Searcher = (*SearcherClient)(nil)

type SearcherClient struct {
	cfg    *config.App
	client *http.Client
	logger logger.Logger
}

func NewSearcherClient(client *http.Client, logger logger.Logger, cfg *config.App) *SearcherClient {
	return &SearcherClient{
		client: client,
		logger: logger,
		cfg:    cfg,
	}
}

func (c *SearcherClient) Search(ctx context.Context, query string) ([]domain.Document, error) {
	const op = "adapter.searcher.client.Search"
	_ = op
	return []domain.Document{}, nil
}
