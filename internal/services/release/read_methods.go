package release

import "github.com/brewbits-co/releasedesk/internal/domains/release"

func (s *service) GetReleaseChannels(appID int) ([]release.Channel, error) {
	channels, err := s.releaseRepo.FindChannelsByAppID(appID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *service) ListReleasesByChannel(appID int, channelID int) ([]release.BasicInfo, error) {
	releases, err := s.releaseRepo.FindByAppIDAndChannel(appID, channelID)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (s *service) GetReleaseSummary(appID int, version string) (release.Release, error) {
	releaseEntity, err := s.releaseRepo.GetByAppIDAndVersion(appID, version)
	if err != nil {
		return release.Release{}, err
	}

	return releaseEntity, nil
}

// GetReleaseByID retrieves a release by its ID
func (s *service) GetReleaseByID(releaseID int) (release.Release, error) {
	releaseEntity, err := s.releaseRepo.GetByID(releaseID)
	if err != nil {
		return release.Release{}, err
	}

	return releaseEntity, nil
}
