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
	Snippet    string    `json:"snippet"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type UploadedDocument struct {
	DocumentID string
	FileName   string
	UploadedAt time.Time
}

type DocumentListQuery struct {
	Query     string
	Paginator Paginator
}

type Searcher interface {
	Search(ctx context.Context, query DocumentListQuery) ([]Document, bool, error)
}

type DocumentService interface {
	GetList(ctx context.Context, query DocumentListQuery) ([]Document, bool, error)
	Upload(ctx context.Context, requestID string, file *multipart.FileHeader) (UploadedDocument, error)
}
