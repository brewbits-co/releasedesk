package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"sort"
)

func (s *service) GetUserAccessibleProducts(userID int) ([]product.Product, error) {
	products, err := s.productRepo.Find()
	if err != nil {
		return nil, err
	}

	for i := range products {
		if s.productRepo.GetPlatformAvailability(&products[i]) != nil {
			return nil, err
		}
	}

	return products, nil
}

func (s *service) GetProductOverview(slug values.Slug) (product.Overview, error) {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return product.Overview{}, err
	}

	data := product.Overview{
		SetupGuideCompleted: productEntity.SetupGuideCompleted,
	}

	return data, nil
}

func (s *service) GetCurrentProductData(slug values.Slug) (session.CurrentProductData, error) {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return session.CurrentProductData{}, err
	}

	apps, err := s.appRepo.FindByProductID(productEntity.ID)
	if err != nil {
		return session.CurrentProductData{}, err
	}

	var productApps []session.CurrentProductAppData
	for _, appEntity := range apps {
		productApps = append(productApps, session.CurrentProductAppData{
			AppID:       appEntity.ID,
			AppName:     appEntity.Name,
			AppPlatform: appEntity.Platform,
			ProductSlug: slug,
		})
	}

	// Desired platform order
	platformOrder := map[values.Platform]int{
		values.Android: 1,
		values.IOS:     2,
		values.Windows: 3,
		values.MacOS:   4,
		values.Linux:   5,
	}

	// Sorting based on the platform order
	sort.Slice(productApps, func(i, j int) bool {
		return platformOrder[productApps[i].AppPlatform] < platformOrder[productApps[j].AppPlatform]
	})

	return session.CurrentProductData{
		ProductID:   productEntity.ID,
		ProductName: productEntity.Name,
		ProductSlug: productEntity.Slug,
		Apps:        productApps,
	}, nil
}
