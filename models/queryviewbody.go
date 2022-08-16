package models

type QueryViewBody struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Views      *struct {
		View []View `json:"view,omitempty"`
	} `json:"views,omitempty"`
}
