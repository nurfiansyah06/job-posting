package helper

import "math"

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func NewPagination(page, totalRows, limit int) *Pagination {
	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
	if page < 1 {
		page = 1
	}

	return &Pagination{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}
}
