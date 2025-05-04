package platform

import "errors"

var (
	ErrPlatformNotFound = errors.New("the platform was not found")
	ErrEmptyField       = errors.New("a required field is empty")
)
