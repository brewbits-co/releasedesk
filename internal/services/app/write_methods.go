package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) CreateApp(slug values.Slug, info app.BasicInfo) (app.Platform, error) {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return app.Platform{}, product.ErrProductNotFound
	}

	info.ProductID = productEntity.ID
	newApp := app.NewApp(info)

	err = s.appRepo.Save(&newApp)
	if err != nil {
		return app.Platform{}, err
	}

	return newApp, nil
}
