package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
)

func (s *service) CreateApp(info product.BasicInfo) (product.App, error) {
	info.SetupGuideCompleted = false
	info.Slug.Format()

	newApp := product.NewApp(info)

	err := newApp.IsValid()
	if err != nil {
		return product.App{}, err
	}

	err = s.appRepo.Save(&newApp)
	if err != nil {
		return product.App{}, err
	}

	return newApp, nil
}
