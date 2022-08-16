package tableau

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"time"
)

const (
	DefaultVersion = "3.10"

	filterByNameIn             = `&filter=name:in:[%s]`
	pagingParams               = `%s?pageSize=%d&pageNumber=%d%s`
	mapAssetsParams            = `%s?mapAssetsTo=%s`
	downloadPDFParams          = `%s?type=A4&orientation=Portrait&maxAge=%d`
	queryViewImageParams       = `%s?resolution=high&maxAge=%d`
	queryViewPDFParams         = `%s?type=A4&orientation=Portrait&maxAge=%d`
	getViewByPathParams        = `%s?pageSize=%d&pageNumber=%d&filter=viewUrlName:eq:%s`
	queryViewForSiteParams     = `%s?pageSize=%d&pageNumber=%d`
	queryViewForWorkbookParams = `%s?pageSize=%d&pageNumber=%d`
	signInUri                  = `auth/signin`
	signOutUri                 = `auth/signout`
	switchSiteUri              = `auth/switchSite`
	addUserToGroupUri          = `sites/%s/groups/%s/users`
	addUserToSiteUri           = `sites/%s/users`
	createGroupUri             = `sites/%s/groups`
	deleteGroupUri             = `sites/%s/groups/%s`
	getGroupsForUserUri        = `sites/%s/users/%s/groups`
	getUsersInGroupUri         = `sites/%s/groups/%s/users`
	getUsersOnSiteUri          = `sites/%s/users`
	queryGroupsUri             = `sites/%s/groups`
	queryUserOnSiteUri         = `sites/%s/users/%s`
	removeUserFromSiteUri      = `sites/%s/users/%s`
	removeUserFromGroupUri     = `sites/%s/groups/%s/users/%s`
	updateGroupUri             = `sites/%s/groups/%s`
	updateUserUri              = `sites/%s/users/%s`
	addTagsToViewUri           = `sites/%s/views/%s/tags`
	addTagsToWorkbookUri       = `sites/%s/workbooks/%s/tags`
	deleteTagFromViewUri       = `sites/%s/views/%s/tags/%s`
	deleteTagFromWorkbookUri   = `sites/%s/workbooks/%s/tags/%s`
	downloadWorkbookPDFUri     = `sites/%s/workbooks/%s/pdf`
	getViewUri                 = `sites/%s/views/%s`
	getViewByPathUri           = `sites/%s/views`
	queryViewsForSiteUri       = `sites/%s/views`
	queryViewsForWorkbookUri   = `sites/%s/workbooks/%s/views`
	queryViewImageUri          = `sites/%s/views/%s/image`
	queryViewPDFUri            = `sites/%s/views/%s/pdf`
	queryWorkbookUri           = `sites/%s/workbooks/%s`

	tokenLifetime = 120 * time.Minute
	pageSize      = 500
	defaultMaxAge = 60

	contentTypeHeader   = `Content-Type`
	acceptHeader        = `Accept`
	mimeTypeJSON        = `application/json`
	mimeTypeAny         = `*/*`
	authorizationHeader = `Authorization`
	bearerAuthorization = `Bearer %v`
)

var (
	// NOTE: Used for unused function error
	_ = NewClient

	json = jsoniter.ConfigCompatibleWithStandardLibrary

	ErrInvalidHost                 = errors.New("not a valid host")
	ErrInvalidUsernamePassword     = errors.New("not a valid username or password")
	ErrFailedUnmarshalResponseBody = errors.New("failed to unmarshal response body")
	ErrUnknownError                = errors.New("unknown error")

	ErrBadRequest             = errors.New("the content of the request body is missing or incomplete")
	ErrInvalidPageNumber      = errors.New("invalid page number")
	ErrInvalidPageSize        = errors.New("invalid page size")
	ErrInvalidSiteRole        = errors.New("invalid site role")
	ErrMalformedImportElement = errors.New("malformed import element")
	ErrDeleteFailed           = errors.New("delete failed")
	ErrAddTagsWorkbook        = errors.New("add tags to workbook failed")
	ErrDeleteTagFromWorkbook  = errors.New("delete tag from workbook failed")
	ErrAddTagsView            = errors.New("add tags to view failed")
	ErrDeleteTagFromView      = errors.New("delete tag from view failed")
	ErrDownloadWorkbookPDF    = errors.New("failed to download workbook as PDF")

	ErrNoCredential      = errors.New("no credentials were provided")
	ErrLoginError        = errors.New("the credentials are invalid (wrong username/password) or blocked")
	ErrInvalidCredential = errors.New("invalid credentials were provided")
	ErrSwitchSiteError   = errors.New("cannot switch site; the site might be unavailable or was not found")

	ErrForbidden                    = errors.New("user do not have sufficient permissions")
	ErrActiveDirectoryNotConfigured = errors.New("active directory was not configured")
	ErrPageSizeExceeded             = errors.New("the specified page size in larger than maximum page size")
	ErrImportNameForbidden          = errors.New("imported name element different with referenced group-id")
	ErrReadForbidden                = errors.New("do not have read access to this resource")
	ErrCannotSwitchToSameSite       = errors.New("cannot switch to the same site")
	ErrDownloadPDFDisabled          = errors.New("download PDF was disabled")
	ErrQueryUserForbidden           = errors.New("user does not have permission to query user information")

	ErrSiteNotFound                 = errors.New("site was not found")
	ErrVersionNotFound              = errors.New("invalid version were provided")
	ErrUserNotFound                 = errors.New("user was not found")
	ErrWorkbookNotFound             = errors.New("workbook was not found")
	ErrTagNotFound                  = errors.New("tag was not found")
	ErrWorkbookIDMismatch           = errors.New("workbook id mismatch")
	ErrViewNotFound                 = errors.New("view was not found")
	ErrGroupNotFound                = errors.New("group was not found")
	ErrDomainNotFound               = errors.New("domain was not found")
	ErrActiveDirectoryGroupNotFound = errors.New("active directory group was not found")

	ErrInvalidRequestMethod = errors.New("not a valid request type")

	ErrUserAlreadyOnSite      = errors.New("the specified user already exist on the site")
	ErrUserAssetConflict      = errors.New("user still owns content and cannot be deleted")
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
		"400032": ErrDeleteFailed,
		"400049": ErrAddTagsWorkbook,
		"400051": ErrDeleteTagFromWorkbook,
		"400076": ErrAddTagsView,
		"400078": ErrDeleteTagFromView,
		"400101": ErrDownloadWorkbookPDF,

		"401000": ErrNoCredential,
		"401001": ErrLoginError,
		"401002": ErrInvalidCredential,
		"401003": ErrSwitchSiteError,

		"403004": ErrForbidden,
		"403011": ErrActiveDirectoryNotConfigured,
		"403014": ErrPageSizeExceeded,
		"403020": ErrImportNameForbidden,
		"403032": ErrReadForbidden,
		"403070": ErrCannotSwitchToSameSite,
		"403105": ErrDownloadPDFDisabled,
		"403133": ErrQueryUserForbidden,

		"404000": ErrSiteNotFound,
		"404001": ErrVersionNotFound,
		"404002": ErrUserNotFound,
		"404006": ErrWorkbookNotFound,
		"404007": ErrTagNotFound,
		"404009": ErrWorkbookIDMismatch,
		"404011": ErrViewNotFound,
		"404012": ErrGroupNotFound,
		"404016": ErrDomainNotFound,
		"404017": ErrActiveDirectoryGroupNotFound,

		"405000": ErrInvalidRequestMethod,

		"409000": ErrUserAlreadyOnSite,
		"409003": ErrUserAssetConflict,
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
