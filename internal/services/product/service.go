package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

// Service defines the interface for handling application-related use cases.
type Service interface {
	// CreateApp creates a new application using the provided basic information.
	// It validates the input data against defined business rules and returns
	// the created application or an error if the validation or creation fails.
	CreateApp(info product.BasicInfo) (product.App, error)
	// GetUserAccessibleApps retrieves the list of applications that a given user
	// has access to, based on the user's ID.
	GetUserAccessibleApps(userID int) ([]product.App, error)
	// ApplyAppSetupGuide applies the setup guide for the specified application,
	// configuring it with the chosen settings.
	ApplyAppSetupGuide(slug values.Slug, format values.VersionFormat, channels product.SetupChannelsOption, customChannels []string) error
	// GetAppOverview returns a summary of an application.
	GetAppOverview(slug values.Slug) (product.Overview, error)
	// GetCurrentAppData retrieves application information shared across views based on the provided slug.
	GetCurrentAppData(slug values.Slug) (session.CurrentAppData, error)
}

// NewAppService initializes a new instance of the application Service using the provided dependencies.
func NewAppService(appRepo product.AppRepository, platformRepo platform.PlatformRepository) Service {
	return &service{
		appRepo:      appRepo,
		platformRepo: platformRepo,
	}
}

// service implements the product.Service
type service struct {
	appRepo      product.AppRepository
	platformRepo platform.PlatformRepository
}
