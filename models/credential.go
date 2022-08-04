package models

type Credentials struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Site     *Site  `json:"site,omitempty"`
	User     *User  `json:"user,omitempty"`
	Token    string `json:"token,omitempty"`
}
