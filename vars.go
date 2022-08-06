package tableau

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"time"
)

const (
	DefaultVersion = "3.10"

	filterByNameIn          = `&filter=name:in:[%s]`
	pagingParams            = `%s?pageSize=%d&pageNumber=%d%s`
	mapAssetsParams         = `%s?mapAssetsTo=%s`
	signInPath              = `auth/signin`
	signOutPath             = `auth/signout`
	switchSitePath          = `auth/switchSite`
	addUserToGroupPath      = `sites/%s/groups/%s/users`
	addUserToSitePath       = `sites/%s/users`
	createGroupPath         = `sites/%s/groups`
	deleteGroupPath         = `sites/%s/groups/%s`
	getGroupsForUserPath    = `sites/%s/users/%s/groups`
	getUsersInGroupPath     = `sites/%s/groups/%s/users`
	getUsersOnSitePath      = `sites/%s/users`
	queryGroupsPath         = `sites/%s/groups`
	queryUserOnSitePath     = `sites/%s/users/%s`
	removeUserFromSitePath  = `sites/%s/users/%s`
	removeUserFromGroupPath = `sites/%s/groups/%s/users/%s`
	updateGroupPath         = `sites/%s/groups/%s`
	updateUserPath          = `sites/%s/users/%s`

	tokenLifetime = 120 * time.Minute
	pageSize      = 500

	contentTypeHeader   = "Content-Type"
	acceptHeader        = "Accept"
	mimeTypeJson        = "application/json"
	authorizationHeader = "Authorization"
	bearerAuthorization = "Bearer %v"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	ErrInvalidHost                 = errors.New("not a valid host")
	ErrInvalidUsernamePassword     = errors.New("not a valid username or password")
	ErrUnknownError                = errors.New("unknown error")
	ErrFailedUnmarshalResponseBody = errors.New("failed to unmarshal response body")

	ErrBadRequest             = errors.New("the content of the request body is missing or incomplete")
	ErrInvalidPageNumber      = errors.New("invalid page number")
	ErrInvalidPageSize        = errors.New("invalid page size")
	ErrInvalidSiteRole        = errors.New("invalid site role")
	ErrMalformedImportElement = errors.New("malformed import element")

	ErrNoCredential      = errors.New("no credentials were provided")
	ErrLoginError        = errors.New("the credentials are invalid (wrong username/password) or blocked")
	ErrInvalidCredential = errors.New("invalid credentials were provided")
	ErrSwitchSiteError   = errors.New("cannot switch site; the site might be unavailable or was not found")

	ErrForbidden                    = errors.New("user do not have sufficient permissions")
	ErrActiveDirectoryNotConfigured = errors.New("active directory was not configured")
	ErrPageSizeExceeded             = errors.New("the specified page size in larger than maximum page size")
	ErrCannotSwitchToSameSite       = errors.New("cannot switch to the same site")

	ErrSiteNotFound                 = errors.New("site was not found")
	ErrVersionNotFound              = errors.New("invalid version were provided")
	ErrUserNotFound                 = errors.New("user was not found")
	ErrGroupNotFound                = errors.New("group was not found")
	ErrDomainNotFound               = errors.New("domain was not found")
	ErrActiveDirectoryGroupNotFound = errors.New("active directory group was not found")

	ErrInvalidRequestMethod = errors.New("not a valid request type")

	ErrUserAlreadyOnSite      = errors.New("the specified user already exist on the site")
	ErrGuestUserNotAllowed    = errors.New("adding user to a site with guest role was not allowed")
	ErrGroupNameAlreadyExists = errors.New("group name already exists")
	ErrUserAlreadyInGroup     = errors.New("the specified user already a member of the group")

	ErrPayloadToLarge = errors.New("request body is too large to process")

	ErrTooManyRequest = errors.New("request limit reached")

	ErrInternalServerError   = errors.New("tableau server error")
	ErrInternalServiceError  = errors.New("tableau service error")
	ErrBroadcastServiceError = errors.New("broadcast service error")

	errCodeMap = map[string]error{
		"400000": ErrBadRequest,
		"400006": ErrInvalidPageNumber,
		"400007": ErrInvalidPageSize,
		"400013": ErrInvalidSiteRole,
		"400019": ErrMalformedImportElement,
		"401000": ErrNoCredential,
		"401001": ErrLoginError,
		"401002": ErrInvalidCredential,
		"401003": ErrSwitchSiteError,
		"403004": ErrForbidden,
		"403011": ErrActiveDirectoryNotConfigured,
		"403014": ErrPageSizeExceeded,
		"403070": ErrCannotSwitchToSameSite,
		"404000": ErrSiteNotFound,
		"404001": ErrVersionNotFound,
		"404002": ErrUserNotFound,
		"404012": ErrGroupNotFound,
		"404016": ErrDomainNotFound,
		"404017": ErrActiveDirectoryGroupNotFound,
		"405000": ErrInvalidRequestMethod,
		"409000": ErrUserAlreadyOnSite,
		"409005": ErrGuestUserNotAllowed,
		"409009": ErrGroupNameAlreadyExists,
		"409011": ErrUserAlreadyInGroup,
		"413000": ErrPayloadToLarge,
		"429000": ErrTooManyRequest,
		"500000": ErrInternalServerError,
		"500001": ErrInternalServiceError,
		"500002": ErrBroadcastServiceError,
	}
)
