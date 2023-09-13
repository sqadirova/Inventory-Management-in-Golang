package util

import (
	"math"
)

type Pagination struct {
	Limit  int64 `json:"limit" validate:"gte=10"`
	Offset int64 `json:"offset"`
}

func AddPagination(page int, pagination *Pagination, count int64) (*Pagination, int, int) {
	maxPageCount := int(math.Ceil(float64(count) / float64(pagination.Limit)))

	if page > maxPageCount {
		return nil, 0, 0
	}

	if page > 0 {
		pagination.Offset = int64(page-1) * pagination.Limit
	}

	nextPage := page + 1

	if page == maxPageCount {
		nextPage = 0
	}

	return pagination, nextPage, maxPageCount
}
