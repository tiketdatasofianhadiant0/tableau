package models

import "strconv"

type Pagination struct {
	PageNumber     string `json:"pageNumber,omitempty"`
	PageSize       string `json:"pageSize,omitempty"`
	TotalAvailable string `json:"totalAvailable,omitempty"`
}

func (p Pagination) GetPageNumber() int {
	if n, err := strconv.Atoi(p.PageNumber); err != nil {
		return 0
	} else {
		return n
	}
}

func (p Pagination) GetPageSize() int {
	if n, err := strconv.Atoi(p.PageSize); err != nil {
		return 0
	} else {
		return n
	}
}

func (p Pagination) GetTotalAvailable() int {
	if n, err := strconv.Atoi(p.TotalAvailable); err != nil {
		return 0
	} else {
		return n
	}
}
