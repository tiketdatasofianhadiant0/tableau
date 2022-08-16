package models

import "time"

type Workbook struct {
	ID              *string    `json:"id,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	WebpageUrl      *string    `json:"webpageUrl,omitempty"`
	ContentUrl      *string    `json:"contentUrl,omitempty"`
	ShowTabs        *string    `json:"showTabs,omitempty"`
	Size            *string    `json:"size,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	DefaultViewID   *string    `json:"defaultViewId,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	EncryptExtracts *string    `json:"encryptExtracts,omitempty"`
	Project         *Project   `json:"project,omitempty"`
	Owner           *Owner     `json:"owner,omitempty"`
	Tags            *struct {
		Tag []Tag `json:"tag,omitempty"`
	} `json:"tags,omitempty"`
	Views *struct {
		View []View `json:"view,omitempty"`
	} `json:"views,omitempty"`
}
