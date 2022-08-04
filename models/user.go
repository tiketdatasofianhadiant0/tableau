package models

import "time"

type User struct {
	ID                 string     `json:"id,omitempty"`
	ExternalAuthUserID string     `json:"externalAuthUserId,omitempty"`
	Name               string     `json:"name,omitempty"`
	FullName           string     `json:"fullName,omitempty"`
	Email              string     `json:"email,omitempty"`
	Password           string     `json:"password,omitempty"`
	AuthSetting        string     `json:"authSetting,omitempty"`
	LastLogin          *time.Time `json:"lastLogin,omitempty"`
	SiteRole           string     `json:"siteRole,omitempty"`
	Locale             string     `json:"locale,omitempty"`
	Language           string     `json:"language,omitempty"`
	Domain             *Domain    `json:"domain,omitempty"`
}
