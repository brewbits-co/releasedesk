package release

import "github.com/brewbits-co/releasedesk/internal/domains/release"

func (s *service) GetReleaseChannels(productID int) ([]release.Channel, error) {
	channels, err := s.releaseRepo.FindChannelsByProductID(productID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *service) ListReleasesByChannel(productID int, channelID int) ([]release.BasicInfo, error) {
	releases, err := s.releaseRepo.FindByProductIDAndChannel(productID, channelID)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (s *service) GetReleaseSummary(productID int, version string) (release.Release, error) {
	releaseEntity, err := s.releaseRepo.GetByProductIdAndVersion(productID, version)
	if err != nil {
		return release.Release{}, err
	}

	return releaseEntity, nil
}
