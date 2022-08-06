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
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
