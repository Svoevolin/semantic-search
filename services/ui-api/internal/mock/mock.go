package mock

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
)

type MockDocumentService struct{}

func NewMockDocumentService() *MockDocumentService {
	return &MockDocumentService{}
}

func (m *MockDocumentService) GetList(ctx context.Context, query domain.DocumentListQuery) ([]domain.Document, bool, error) {
	now := time.Now()
	return []domain.Document{
		{
			DocumentID: uuid.NewString(),
			FileName:   "semantic-guide.pdf",
			MatchScore: 0.93,
			UploadedAt: now.Add(-2 * time.Hour),
			Snippet:    "Фрагмент документа в котором нашлись совпадения документадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокумента",
		},
		{
			DocumentID: uuid.NewString(),
			FileName:   "ai-whitepaper.pdf",
			MatchScore: 0.89,
			UploadedAt: now.Add(-6 * time.Hour),
			Snippet:    "Фрагмент документа в котором нашлись совпадения 2 документадокументадокументадокументадокументадокументадокументадокументадокумента",
		},
		{
			DocumentID: uuid.NewString(),
			FileName:   "eco-policy.pdf",
			MatchScore: 0.84,
			UploadedAt: now.Add(-24 * time.Hour),
			Snippet:    "Фрагмент документа в котором нашлись совпадения 3 документадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокументадокумента",
		},
	}, false, nil
}

func (m *MockDocumentService) Upload(ctx context.Context, requestID string, file *multipart.FileHeader) (domain.UploadedDocument, error) {
	return domain.UploadedDocument{
		DocumentID: uuid.NewString(),
		FileName:   file.Filename,
		UploadedAt: time.Now(),
	}, nil
}

// -- Mock Storage --

type MockStorage struct{}

func (m *MockStorage) Upload(ctx context.Context, file *multipart.FileHeader) (domain.UploadResult, error) {
	return domain.UploadResult{
		ObjectName: "mock/path/to/" + file.Filename,
		URL:        &url.URL{Scheme: "http", Host: "localhost:9000", Path: "/mock/path/to/" + file.Filename},
	}, nil
}

// -- Mock Kafka --

type MockKafkaProducer struct{}

func (m *MockKafkaProducer) PublishUpload(ctx context.Context, event domain.UploadEvent) error {
	return nil // просто игнорируем
}

// -- Mock Searcher --

type MockSearcherClient struct{}

func (m *MockSearcherClient) Search(ctx context.Context, query string) ([]domain.Document, error) {
	return nil, nil
}
