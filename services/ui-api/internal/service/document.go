package service

import (
	"context"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
)

type ClientSearcher interface {
	Search(ctx context.Context, query string) ([]domain.Document, error)
}

type Document struct {
	client ClientSearcher
}

func NewDocument(client ClientSearcher) *Document {
	return &Document{client: client}
}

func (s *Document) GetList(ctx context.Context, query domain.DocumentListQuery) ([]domain.Document, error) {
	return s.client.Search(ctx, query.Query)
}
