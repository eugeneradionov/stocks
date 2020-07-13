package auth

import "errors"

var (
	ErrEmailExists            = errors.New("email is already exists")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
