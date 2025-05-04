package build

import "github.com/brewbits-co/releasedesk/internal/domains/build"

func (s *service) GetPlatformBuilds(platformID int) ([]build.BasicInfo, error) {
	builds, err := s.buildRepo.FindByPlatformID(platformID)
	if err != nil {
		return nil, err
	}

	return builds, nil
}

func (s *service) GetBuildDetails(platformID int, number int) (build.Build, error) {
	buildEntity, err := s.buildRepo.GetByPlatformIDAndNumber(platformID, number)
	if err != nil {
		return build.Build{}, err
	}

	return buildEntity, err
}
