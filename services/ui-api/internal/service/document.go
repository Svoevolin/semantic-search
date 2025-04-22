package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
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

func (s *Document) Upload(_ context.Context, file *multipart.FileHeader) (domain.UploadedDocument, error) {
	// TODO: Загрузка в MinIO и отправка в Kafka
	// Сейчас просто имитация
	return domain.UploadedDocument{
		DocumentID: uuid.New().String(),
		FileName:   file.Filename,
		UploadedAt: time.Now().UTC(),
	}, nil
}
