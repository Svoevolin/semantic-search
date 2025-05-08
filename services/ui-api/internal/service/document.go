package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
)

type Document struct {
	searcher      domain.Searcher
	storageClient domain.StorageUploader
	producer      domain.KafkaProducer
	logger        logger.Logger
}

func NewDocument(
	searcher domain.Searcher,
	storage domain.StorageUploader,
	producer domain.KafkaProducer,
	logger logger.Logger,
) *Document {
	return &Document{
		searcher:      searcher,
		storageClient: storage,
		producer:      producer,
		logger:        logger,
	}
}

func (s *Document) GetList(ctx context.Context, query domain.DocumentListQuery) ([]domain.Document, bool, error) {
	return s.searcher.Search(ctx, query)
}

func (s *Document) Upload(ctx context.Context, requestID string, file *multipart.FileHeader) (domain.UploadedDocument, error) {
	const op = "service.Document.Upload"

	docID := uuid.New().String()

	uploadRes, err := s.storageClient.Upload(ctx, file)
	if err != nil {
		return domain.UploadedDocument{}, fmt.Errorf("upload to storage failed: %w", err)
	}

	err = s.producer.PublishUpload(ctx, domain.UploadEvent{
		DocumentID: docID,
		FileName:   file.Filename,
		ObjectURL:  uploadRes.URL.String(),
		RequestID:  requestID,
	})
	if err != nil {
		s.logger.ErrorContext(ctx, op, "publish to kafka failed", sl.Err(err))
		return domain.UploadedDocument{}, fmt.Errorf("publish to kafka failed: %w", err)
	}

	return domain.UploadedDocument{
		DocumentID: docID,
		FileName:   file.Filename,
		UploadedAt: time.Now().UTC(),
	}, nil
}
