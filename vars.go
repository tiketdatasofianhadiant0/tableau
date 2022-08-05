package tableau

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

const (
	DefaultVersion = "3.10"

	signInPath         = `auth/signin`
	signOutPath        = `auth/signout`
	switchSitePath     = `auth/switchSite`
	addUserToGroupPath = `sites/%s/groups/%s/users`
	addUserToSitePath  = `sites/%s/users`

	retryCount       = 3
	retryWaitTime    = 100 * time.Millisecond
	retryMaxWaitTime = 2 * time.Second
	tokenLifetime    = 120 * time.Minute
	pageSize         = 500

	contentTypeHeader   = "Content-Type"
	acceptHeader        = "Accept"
	mimeTypeJson        = "application/json"
	authorizationHeader = "Authorization"
	bearerAuthorization = "Bearer %v"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
