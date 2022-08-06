package models

type UserBody struct {
	User   *User   `json:"user,omitempty"`
	Domain *Domain `json:"domain,omitempty"`
}
