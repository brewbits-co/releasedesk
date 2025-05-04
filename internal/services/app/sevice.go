package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
)

// Service defines the interface for handling app-related use cases.
type Service interface {
	// CreateApp creates a new app for a specific product using the provided basic information.
	// It validates the input data against defined business rules and returns
	// the created app or an error if the validation or creation fails.
	CreateApp(slug values.Slug, info app.BasicInfo) (app.Platform, error)
}

// NewAppService initializes a new instance of the app Service using the provided dependencies.
func NewAppService(appRepo app.PlatformRepository, productRepo product.ProductRepository) Service {
	return &service{
		appRepo:     appRepo,
		productRepo: productRepo,
	}
}

// service implements the app.Service
type service struct {
	appRepo     app.PlatformRepository
	productRepo product.ProductRepository
}
