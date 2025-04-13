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
