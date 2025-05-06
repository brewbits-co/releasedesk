package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

// Service defines the interface for handling release-related use cases.
type Service interface {
	CreateRelease(slug values.Slug, info release.BasicInfo) (release.Release, error)
	ListReleasesByChannel(appID int, channelID int) ([]release.BasicInfo, error)
	GetReleaseSummary(appID int, version string) (release.Release, error)
	GetReleaseChannels(appID int) ([]release.Channel, error)
}

// NewReleaseService initializes a new instance of the release Service using the provided dependencies.
func NewReleaseService(releaseRepo release.ReleaseRepository, appRepo app.AppRepository) Service {
	return &service{
		releaseRepo: releaseRepo,
		appRepo:     appRepo,
	}
}

type service struct {
	releaseRepo release.ReleaseRepository
	appRepo     app.AppRepository
}
