package domain

import (
	"context"
	"mime/multipart"
	"time"
)

type Document struct {
	DocumentID string    `json:"document_id"`
	FileName   string    `json:"file_name"`
	MatchScore float64   `json:"match_score"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type UploadedDocument struct {
	DocumentID string
	FileName   string
	UploadedAt time.Time
}

type DocumentListQuery struct {
	Query string
}

type Searcher interface {
	Search(ctx context.Context, query string) ([]Document, error)
}

type DocumentService interface {
	GetList(ctx context.Context, query DocumentListQuery) ([]Document, error)
	Upload(ctx context.Context, file *multipart.FileHeader) (UploadedDocument, error)
}
