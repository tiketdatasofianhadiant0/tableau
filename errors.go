package tableau

import "errors"

var (
	ErrInvalidHost             = errors.New("not a valid host")
	ErrInvalidUsernamePassword = errors.New("not a valid username or password")
)
