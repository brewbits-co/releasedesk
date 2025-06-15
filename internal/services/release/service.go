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
	SaveReleaseNotes(releaseID int, text string, changelogs []release.Changelog) (release.ReleaseNotes, error)
	GetReleaseNotes(releaseID int) (release.ReleaseNotes, error)
}

// NewReleaseService initializes a new instance of the release Service using the provided dependencies.
func NewReleaseService(releaseRepo release.ReleaseRepository, releaseNotesRepo release.ReleaseNotesRepository, appRepo app.AppRepository) Service {
	return &service{
		releaseRepo:      releaseRepo,
		releaseNotesRepo: releaseNotesRepo,
		appRepo:          appRepo,
	}
}

type service struct {
	releaseRepo      release.ReleaseRepository
	releaseNotesRepo release.ReleaseNotesRepository
	appRepo          app.AppRepository
}

// SaveReleaseNotes saves release notes and associated changelogs for a release
func (s *service) SaveReleaseNotes(releaseID int, text string, changelogs []release.Changelog) (release.ReleaseNotes, error) {
	releaseNotes := release.NewReleaseNotes(releaseID, text)
	releaseNotes.Changelogs = changelogs

	if err := s.releaseNotesRepo.Save(&releaseNotes); err != nil {
		return release.ReleaseNotes{}, err
	}

	return releaseNotes, nil
}

// GetReleaseNotes retrieves release notes and associated changelogs for a release
func (s *service) GetReleaseNotes(releaseID int) (release.ReleaseNotes, error) {
	releaseNotes, err := s.releaseNotesRepo.GetByReleaseID(releaseID)
	if err != nil {
		return release.ReleaseNotes{}, err
	}

	return releaseNotes, nil
}
