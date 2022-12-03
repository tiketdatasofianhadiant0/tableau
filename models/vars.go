package models

import jsoniter "github.com/json-iterator/go"

const (
	SiteRoleUnlicensed                = `Unlicensed`
	SiteRoleViewer                    = `Viewer`
	SiteRoleExplorer                  = `Explorer`
	SiteRoleExplorerCanPublish        = `ExplorerCanPublish`
	SiteRoleCreator                   = `Creator`
	SiteRoleSiteAdministratorExplorer = `SiteAdministratorExplorer`
	SiteRoleSiteAdministratorCreator  = `SiteAdministratorCreator`

	ImageResolutionHigh = `high`

	defaultMaxAge = 60
	minMaxAge     = 1
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
