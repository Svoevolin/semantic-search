package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/delivery/http/dto"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type Document struct {
	service domain.DocumentService
	logger  logger.Logger
}

func NewDocument(service domain.DocumentService, logger logger.Logger) *Document {
	return &Document{service: service, logger: logger}
}

// swagger:route POST /documents documents listDocuments
//
// # Semantic document search
//
// Performs a semantic search over available documents using the given query string.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Responses:
//
//	200: DocumentListResponse
//	400: ErrorResponse
//	500: ErrorResponse
func (h *Document) List(c echo.Context) error {
	const op = "handlers.document.List"
	ctx := c.Request().Context()

	req := dto.DocumentListRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	results, err := h.service.GetList(ctx, req.Model())
	if err != nil {
		return err
	}
	h.logger.InfoContext(ctx, op, "documents found", "count", len(results))

	return c.JSON(http.StatusOK, dto.NewDocumentListResponse(results).Body)
}

// swagger:route POST /documents/upload documents uploadDocument
//
// # Upload a new PDF document
//
// Uploads a machine-readable PDF file to the server.
//
// Consumes:
// - multipart/form-data
//
// Produces:
// - application/json
//
// Responses:
//
//	200: UploadDocumentResponse
//	400: ErrorResponse
//	500: ErrorResponse
func (h *Document) Upload(c echo.Context) error {
	const op = "handlers.document.Upload"
	ctx := c.Request().Context()

	req := dto.UploadDocumentRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	doc := req.Model()
	h.logger.InfoContext(ctx, op, "upload file received", "filename", doc.Filename, "size", doc.Size)

	result, err := h.service.Upload(ctx, doc)
	if err != nil {
		return err
	}
	h.logger.InfoContext(ctx, op, "document uploaded", "document_id", result.DocumentID)

	return c.JSON(http.StatusOK, dto.NewUploadDocumentResponse(result).Body)
}
