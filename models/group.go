package models

type Group struct {
	ID              string  `json:"id,omitempty"`
	Name            string  `json:"name,omitempty"`
	MinimumSiteRole string  `json:"minimumSiteRole,omitempty"`
	Import          *Import `json:"import,omitempty"`
}
