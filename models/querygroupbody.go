package models

type QueryGroupBody struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Groups     *struct {
		Group []Group `json:"group,omitempty"`
	} `json:"groups,omitempty"`
}
