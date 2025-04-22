package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type Document struct {
	searcher      domain.Searcher
	storageClient domain.StorageUploader
	logger        logger.Logger
}

func NewDocument(searcher domain.Searcher, storage domain.StorageUploader, logger logger.Logger) *Document {
	return &Document{
		searcher:      searcher,
		storageClient: storage,
		logger:        logger,
	}
}

func (s *Document) GetList(ctx context.Context, query domain.DocumentListQuery) ([]domain.Document, error) {
	return s.searcher.Search(ctx, query.Query)
}

func (s *Document) Upload(ctx context.Context, file *multipart.FileHeader) (domain.UploadedDocument, error) {
	uploadResult, err := s.storageClient.Upload(ctx, file)
	if err != nil {
		return domain.UploadedDocument{}, fmt.Errorf("upload to storage failed: %w", err)
	}
	_ = uploadResult

	documentID := uuid.New().String()
	_ = documentID

	// TODO: отправить в Kafka: documentID, uploadResult.URL и т.д.

	// Сейчас просто имитация
	return domain.UploadedDocument{
		DocumentID: uuid.New().String(),
		FileName:   file.Filename,
		UploadedAt: time.Now().UTC(),
	}, nil
}
