package tableau

import "errors"

var (
	ErrInvalidHost                 = errors.New("not a valid host")
	ErrInvalidUsernamePassword     = errors.New("not a valid username or password")
	ErrUnknownError                = errors.New("unknown error")
	ErrFailedUnmarshalResponseBody = errors.New("failed to unmarshal response body")

	ErrForbidden              = errors.New("user do not have sufficient permissions")
	ErrServerNotFound         = errors.New("server was not found")
	ErrInvalidRequestMethod   = errors.New("not a valid request type")
	ErrBadRequest             = errors.New("the content of the request body is missing or incomplete")
	ErrLoginError             = errors.New("the credentials are invalid (wrong username/password) or blocked")
	ErrPayloadToLarge         = errors.New("request body is too large to process")
	ErrTooManyRequest         = errors.New("request limit reached")
	ErrNoCredential           = errors.New("no credentials were provided")
	ErrInvalidCredential      = errors.New("invalid credentials were provided")
	ErrSwitchSiteError        = errors.New("cannot switch site; the site might be unavailable or was not found")
	ErrCannotSwitchToSameSite = errors.New("cannot switch to the same site")
	ErrInternalServerError    = errors.New("tableau server error")
	ErrInternalServiceError   = errors.New("tableau service error")
	ErrBroadcastServiceError  = errors.New("broadcast service error")

	errCodeMap = map[string]error{
		"400000": ErrBadRequest,
		"401000": ErrNoCredential,
		"401001": ErrLoginError,
		"401002": ErrInvalidCredential,
		"401003": ErrSwitchSiteError,
		"403004": ErrForbidden,
		"403070": ErrCannotSwitchToSameSite,
		"404000": ErrServerNotFound,
		"405000": ErrInvalidRequestMethod,
		"413000": ErrPayloadToLarge,
		"429000": ErrTooManyRequest,
		"500000": ErrInternalServerError,
		"500001": ErrInternalServiceError,
		"500002": ErrBroadcastServiceError,
	}
)
