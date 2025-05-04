package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) AddPlatformToApp(slug values.Slug, info platform.BasicInfo) (platform.Platform, error) {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return platform.Platform{}, product.ErrProductNotFound
	}

	info.AppID = productEntity.ID
	addedPlatform := platform.NewApp(info)

	err = s.platformRepo.Save(&addedPlatform)
	if err != nil {
		return platform.Platform{}, err
	}

	return addedPlatform, nil
}
