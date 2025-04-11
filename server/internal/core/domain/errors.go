package domain

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired") // TODO: change err msg
	ErrOAuthStateNotFound = errors.New("state not found")
	ErrUserNotFound       = errors.New("user not found")
)
