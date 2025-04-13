package build

import "github.com/brewbits-co/releasedesk/internal/domains/build"

func (s *service) GetAppBuilds(appID int) ([]build.BasicInfo, error) {
	builds, err := s.buildRepo.FindByAppID(appID)
	if err != nil {
		return nil, err
	}

	return builds, nil
}

func (s *service) GetBuildDetails(appID int, number int) (build.Build, error) {
	buildEntity, err := s.buildRepo.GetByAppIDAndNumber(appID, number)
	if err != nil {
		return build.Build{}, err
	}

	return buildEntity, err
}
