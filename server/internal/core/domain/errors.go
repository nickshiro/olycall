package domain

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrOAuthStateNotFound = errors.New("state not found")
	ErrUserNotFound       = errors.New("user not found")
)
