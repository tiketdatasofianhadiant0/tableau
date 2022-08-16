package models

import "time"

type View struct {
	ID          *string    `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	ContentUrl  *string    `json:"contentUrl,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	ViewUrlName *string    `json:"viewUrlName,omitempty"`
	Workbook    *Workbook  `json:"workbook,omitempty"`
	Owner       *Owner     `json:"owner,omitempty"`
	Project     *Project   `json:"project,omitempty"`
	Usage       *Usage     `json:"usage,omitempty"`
}
