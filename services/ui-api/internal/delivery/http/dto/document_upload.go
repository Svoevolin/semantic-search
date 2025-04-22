package dto

import (
	"mime/multipart"
	"time"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
)

// swagger:parameters uploadDocument
type UploadDocumentRequest struct {
	// Uploaded PDF file
	// in: formData
	// required: true
	// swagger:file
	File *multipart.FileHeader `form:"file"`
}

func (r *UploadDocumentRequest) Model() *multipart.FileHeader {
	return r.File
}

// swagger:response UploadDocumentResponse
type UploadDocumentResponse struct {
	// in: body
	Body UploadDocumentResponseBody
}

// swagger:model UploadDocumentResponseBody
type UploadDocumentResponseBody struct {
	// UUID of the uploaded document
	// example: 550e8400-e29b-41d4-a716-446655440000
	DocumentID string `json:"document_id"`

	// Name of the uploaded file
	// example: report.pdf
	FileName string `json:"file_name"`

	// Upload timestamp
	// example: 2025-04-20T10:00:00Z
	UploadedAt time.Time `json:"uploaded_at"`
}

func NewUploadDocumentResponse(d domain.UploadedDocument) UploadDocumentResponse {
	return UploadDocumentResponse{
		Body: UploadDocumentResponseBody{
			DocumentID: d.DocumentID,
			FileName:   d.FileName,
			UploadedAt: d.UploadedAt,
		},
	}
}
