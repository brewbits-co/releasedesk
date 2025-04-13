package user

import "errors"

var (
	ErrInvalidEmail     = errors.New("email address not valid")
	ErrWrongCredentials = errors.New("email or password are incorrect")
	ErrEmptyField       = errors.New("a required field is empty")
	ErrMustHaveRole     = errors.New("a user must have a defined role")
)
