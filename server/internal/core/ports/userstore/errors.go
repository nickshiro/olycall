package userstore

import "errors"

var (
	ErrEmailConflict    = errors.New("email conflict")    // FIXME
	ErrUsernameConflict = errors.New("username conflict") // FIXME
)
