package product

import "errors"

var (
	ErrProductNotFound            = errors.New("the product was not found")
	ErrSetupGuideAlreadyCompleted = errors.New("the setup guide was already completed")
)
