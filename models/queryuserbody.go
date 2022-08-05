package models

type QueryUserBody struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Users      *struct {
		User []User `json:"user,omitempty"`
	} `json:"users,omitempty"`
}
