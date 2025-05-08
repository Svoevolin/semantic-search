package dto

import (
	"time"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
)

// swagger:model DocumentListRequest
type DocumentListRequest struct {
	// Semantic query text
	// example: sustainable development in corporate policy
	Query string `json:"query" validate:"required"`

	Paginator Paginator
}

func (r DocumentListRequest) Model() domain.DocumentListQuery {
	return domain.DocumentListQuery{
		Query:     r.Query,
		Paginator: r.Paginator.Model(),
	}
}

// swagger:response DocumentListResponse
type DocumentListResponse struct {
	// in: body
	// List of matching documents
	Body []DocumentListResponseItem
}

func NewDocumentListResponse(list []domain.Document) DocumentListResponse {
	items := make([]DocumentListResponseItem, 0, len(list))
	for _, d := range list {
		items = append(items, newDocumentListResponseItem(d))
	}
	return DocumentListResponse{
		Body: items,
	}
}

// DocumentListResponseItem - represents a single document in the list.
// swagger:model DocumentListResponseItem
type DocumentListResponseItem struct {
	// UUID of the document
	// example: 550e8400-e29b-41d4-a716-446655440000
	DocumentID string `json:"document_id"`

	// Name of the file
	// example: policy.pdf
	FileName string `json:"file_name"`

	// Similarity score (if applicable)
	// example: 0.91
	MatchScore float64 `json:"match_score"`

	// Upload timestamp
	// example: 2025-04-20T10:00:00Z
	UploadedAt time.Time `json:"uploaded_at"`
}

func newDocumentListResponseItem(d domain.Document) DocumentListResponseItem {
	return DocumentListResponseItem{
		DocumentID: d.DocumentID,
		FileName:   d.FileName,
		MatchScore: d.MatchScore,
		UploadedAt: d.UploadedAt,
	}
}
