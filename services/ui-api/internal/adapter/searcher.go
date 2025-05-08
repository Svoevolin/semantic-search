package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

// compile-time check
var _ domain.Searcher = (*SearcherClient)(nil)

const search = "/api/v1/search"

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

func (c *SearcherClient) Search(ctx context.Context, query domain.DocumentListQuery) ([]domain.Document, bool, error) {
	const op = "adapter.searcher.client.Search"

	// Сформировать тело запроса
	reqBody := map[string]any{
		"query":     query.Query,
		"_page":     query.Paginator.GetPage(),
		"_pagesize": query.Paginator.GetLimit(),
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, false, fmt.Errorf("%s: marshal error: %w", op, err)
	}

	c.logger.DebugContext(ctx, op, "request body", "json", string(body))

	searchURL := c.cfg.URL.ResolveReference(&url.URL{Path: search}).String()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, searchURL, bytes.NewReader(body))
	if err != nil {
		return nil, false, fmt.Errorf("%s: build request: %w", op, err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	c.logger.DebugContext(ctx, op, "sending request", "url", searchURL, "query", query.Query)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, false, fmt.Errorf("%s: request failed: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("%s: unexpected status: %d", op, resp.StatusCode)
	}

	var response struct {
		Results []domain.Document `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, false, fmt.Errorf("%s: decode error: %w", op, err)
	}

	hasMore := strings.ToLower(resp.Header.Get("X-Has-More")) == "true"

	c.logger.InfoContext(ctx, op, "documents received", "count", len(response.Results), "has_more", hasMore)

	return response.Results, hasMore, nil
}
