package platform

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
)

// Service defines the interface for handling platform-related use cases.
type Service interface {
	// AddPlatformToApp adds a new platform for a specific application using the provided basic information.
	// It validates the input data against defined business rules and returns
	// the created platform or an error if the validation or creation fails.
	AddPlatformToApp(slug values.Slug, info platform.BasicInfo) (platform.Platform, error)
}

// NewPlatformService initializes a new instance of the platform Service using the provided dependencies.
func NewPlatformService(platformRepo platform.PlatformRepository, productRepo product.AppRepository) Service {
	return &service{
		platformRepo: platformRepo,
		productRepo:  productRepo,
	}
}

// service implements the platform.Service
type service struct {
	platformRepo platform.PlatformRepository
	productRepo  product.AppRepository
}
