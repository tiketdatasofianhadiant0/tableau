package tableau

import (
	jsoniter "github.com/json-iterator/go"
	"time"
)

const (
	DefaultVersion = "3.10"

	signInPath  = `auth/signin`
	signOutPath = `auth/signout`

	retryCount       = 3
	retryWaitTime    = 100 * time.Millisecond
	retryMaxWaitTime = 2 * time.Second
	tokenLifetime    = 120 * time.Minute

	contentTypeHeader   = "Content-Type"
	acceptHeader        = "Accept"
	mimeTypeJson        = "application/json"
	authorizationHeader = "Authorization"
	bearerAuthorization = "Bearer %v"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)
