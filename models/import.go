package models

type Import struct {
	Source           *string `json:"source,omitempty"`
	DomainName       *string `json:"domainName,omitempty"`
	SiteRole         *string `json:"siteRole,omitempty"`
	GrantLicenseMode *string `json:"grantLicenseMode,omitempty"`
}
