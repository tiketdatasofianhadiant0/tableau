package tableau

import "errors"

var (
	ErrInvalidHost                 = errors.New("not a valid host")
	ErrInvalidUsernamePassword     = errors.New("not a valid username or password")
	ErrUnknownError                = errors.New("unknown error")
	ErrFailedUnmarshalResponseBody = errors.New("failed to unmarshal response body")

	ErrBadRequest             = errors.New("the content of the request body is missing or incomplete")
	ErrInvalidSiteRole        = errors.New("invalid site role")
	ErrMalformedImportElement = errors.New("malformed import element")

	ErrNoCredential      = errors.New("no credentials were provided")
	ErrLoginError        = errors.New("the credentials are invalid (wrong username/password) or blocked")
	ErrInvalidCredential = errors.New("invalid credentials were provided")
	ErrSwitchSiteError   = errors.New("cannot switch site; the site might be unavailable or was not found")

	ErrForbidden                    = errors.New("user do not have sufficient permissions")
	ErrActiveDirectoryNotConfigured = errors.New("active directory was not configured")
	ErrCannotSwitchToSameSite       = errors.New("cannot switch to the same site")

	ErrSiteNotFound                 = errors.New("site was not found")
	ErrVersionNotFound              = errors.New("invalid version were provided")
	ErrUserNotFound                 = errors.New("user was not found")
	ErrGroupNotFound                = errors.New("group was not found")
	ErrDomainNotFound               = errors.New("domain was not found")
	ErrActiveDirectoryGroupNotFound = errors.New("active directory group was not found")

	ErrInvalidRequestMethod = errors.New("not a valid request type")

	ErrGroupNameAlreadyExists = errors.New("group name already exists")
	ErrUserAlreadyInGroup     = errors.New("the specified user already a member of the group")

	ErrPayloadToLarge = errors.New("request body is too large to process")

	ErrTooManyRequest = errors.New("request limit reached")

	ErrInternalServerError   = errors.New("tableau server error")
	ErrInternalServiceError  = errors.New("tableau service error")
	ErrBroadcastServiceError = errors.New("broadcast service error")

	errCodeMap = map[string]error{
		"400000": ErrBadRequest,
		"400013": ErrInvalidSiteRole,
		"400019": ErrMalformedImportElement,
		"401000": ErrNoCredential,
		"401001": ErrLoginError,
		"401002": ErrInvalidCredential,
		"401003": ErrSwitchSiteError,
		"403004": ErrForbidden,
		"403011": ErrActiveDirectoryNotConfigured,
		"403070": ErrCannotSwitchToSameSite,
		"404000": ErrSiteNotFound,
		"404001": ErrVersionNotFound,
		"404002": ErrUserNotFound,
		"404012": ErrGroupNotFound,
		"404016": ErrDomainNotFound,
		"404017": ErrActiveDirectoryGroupNotFound,
		"405000": ErrInvalidRequestMethod,
		"409009": ErrGroupNameAlreadyExists,
		"409011": ErrUserAlreadyInGroup,
		"413000": ErrPayloadToLarge,
		"429000": ErrTooManyRequest,
		"500000": ErrInternalServerError,
		"500001": ErrInternalServiceError,
		"500002": ErrBroadcastServiceError,
	}
)
