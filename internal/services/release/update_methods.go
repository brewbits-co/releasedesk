package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
)

// UpdateReleaseBasicInfo updates the BasicInfo of a release
// The ID, AppID, and Version fields cannot be updated.
func (s *service) UpdateReleaseBasicInfo(appID int, version string, info release.BasicInfo) (release.Release, error) {
	// Get the existing release
	existingRelease, err := s.releaseRepo.GetByAppIDAndVersion(appID, version)
	if err != nil {
		return release.Release{}, err
	}

	// Update only the fields that are allowed to be updated
	existingRelease.TargetChannel = info.TargetChannel
	existingRelease.Status = info.Status

	// Save the updated release
	err = s.releaseRepo.Update(&existingRelease)
	if err != nil {
		return release.Release{}, err
	}

	return existingRelease, nil
}
