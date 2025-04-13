package app

import "errors"

var (
	ErrAppNotFound = errors.New("the app was not found")
	ErrEmptyField  = errors.New("a required field is empty")
)
