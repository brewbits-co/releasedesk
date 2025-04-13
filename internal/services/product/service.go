package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

// Service defines the interface for handling product-related use cases.
type Service interface {
	// CreateProduct creates a new product using the provided basic information.
	// It validates the input data against defined business rules and returns
	// the created product or an error if the validation or creation fails.
	CreateProduct(info product.BasicInfo) (product.Product, error)
	// GetUserAccessibleProducts retrieves the list of products that a given user
	// has access to, based on the user's ID.
	GetUserAccessibleProducts(userID int) ([]product.Product, error)
	// ApplyProductSetupGuide applies the setup guide for the specified product,
	// configuring it with the chosen settings.
	ApplyProductSetupGuide(slug values.Slug, format values.VersionFormat, channels product.SetupChannelsOption, customChannels []string) error
	// GetProductOverview returns a summary of a product.
	GetProductOverview(slug values.Slug) (product.Overview, error)
	// GetCurrentProductData retrieves product information shared across views based on the provided slug.
	GetCurrentProductData(slug values.Slug) (session.CurrentProductData, error)
}

// NewProductService initializes a new instance of the product Service using the provided dependencies.
func NewProductService(productRepo product.ProductRepository, appRepo app.AppRepository) Service {
	return &service{
		productRepo: productRepo,
		appRepo:     appRepo,
	}
}

// service implements the product.Service
type service struct {
	productRepo product.ProductRepository
	appRepo     app.AppRepository
}
