package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) CreateRelease(slug values.Slug, info release.BasicInfo) (release.Release, error) {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return release.Release{}, err
	}

	info.ProductID = productEntity.ID
	info.Status = release.Unpublished

	newRelease := release.NewRelease(info)

	err = s.releaseRepo.Save(&newRelease)
	if err != nil {
		return release.Release{}, err
	}

	return newRelease, nil
}
