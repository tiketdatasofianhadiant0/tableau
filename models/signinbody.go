package models

type SignInBody struct {
	Credentials *Credentials `json:"credentials,omitempty"`
}
