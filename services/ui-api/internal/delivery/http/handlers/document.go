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
	req := dto.DocumentListRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	results, err := h.service.GetList(c.Request().Context(), req.Model())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.NewDocumentListResponse(results).Body)
}
