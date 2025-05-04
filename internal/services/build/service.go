package build

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/values"
	"mime/multipart"
)

// Service defines the interface for handling build-related use cases.
type Service interface {
	// UploadBuild uploads a new build for the specified app, using the provided build information and files.
	UploadBuild(slug values.Slug, platform values.OS, info build.BasicInfo, files map[values.Architecture]*multipart.FileHeader, Metadata map[string]string) (build.Build, error)
	// GetAppBuilds retrieves the list of build from a specific OS.
	GetAppBuilds(appID int) ([]build.BasicInfo, error)
	// GetBuildDetails return the full information of a Build including its Artifacts and Metadata.
	GetBuildDetails(appID int, number int) (build.Build, error)
}

// NewBuildService initializes a new instance of the build Service using the provided dependencies.
func NewBuildService(buildRepo build.BuildRepository, appRepo app.PlatformRepository) Service {
	return &service{
		buildRepo: buildRepo,
		appRepo:   appRepo,
	}
}

// service implements the build.Service
type service struct {
	buildRepo build.BuildRepository
	appRepo   app.PlatformRepository
}
