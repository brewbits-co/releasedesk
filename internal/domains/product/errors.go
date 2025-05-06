package product

import "errors"

var (
	ErrAppNotFound                = errors.New("the application was not found")
	ErrSetupGuideAlreadyCompleted = errors.New("the setup guide was already completed")
)
