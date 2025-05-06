package platform

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) AddPlatformToApp(slug values.Slug, info platform.BasicInfo) (platform.Platform, error) {
	appEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return platform.Platform{}, product.ErrAppNotFound
	}

	info.AppID = appEntity.ID
	addedPlatform := platform.NewPlatform(info)

	err = s.platformRepo.Save(&addedPlatform)
	if err != nil {
		return platform.Platform{}, err
	}

	return addedPlatform, nil
}
