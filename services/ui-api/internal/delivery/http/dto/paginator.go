package dto

import (
	"strconv"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
)

type Paginator struct {
	// Page number (1-based)
	// Example: 1
	Page string `json:"_page" query:"_page" validate:"omitempty,number_gte=1"`

	// Page size (number of items per page)
	// Minimum: 1
	// Maximum: 100
	// Example: 10
	PageSize string `json:"_pagesize" query:"_pagesize" validate:"omitempty,numeric,number_gte=1,number_lte=100"`
}

func (p Paginator) Model() domain.Paginator {
	page, _ := strconv.Atoi(p.Page)
	pagesize, _ := strconv.Atoi(p.PageSize)

	return domain.Paginator{
		Page:     page,
		PageSize: pagesize,
	}
}
