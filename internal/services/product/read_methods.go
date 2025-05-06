package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"sort"
)

func (s *service) GetUserAccessibleApps(userID int) ([]product.App, error) {
	applications, err := s.appRepo.Find()
	if err != nil {
		return nil, err
	}

	for i := range applications {
		if s.appRepo.GetPlatformAvailability(&applications[i]) != nil {
			return nil, err
		}
	}

	return applications, nil
}

func (s *service) GetAppOverview(slug values.Slug) (product.Overview, error) {
	applicationEntity, err := s.appRepo.FindBySlug(slug)
	if err != nil {
		return product.Overview{}, err
	}

	data := product.Overview{
		SetupGuideCompleted: applicationEntity.SetupGuideCompleted,
	}

	return data, nil
}

func (s *service) GetCurrentAppData(slug values.Slug) (session.CurrentAppData, error) {
	applicationEntity, err := s.appRepo.FindBySlug(slug)
	if err != nil {
		return session.CurrentAppData{}, err
	}

	platforms, err := s.platformRepo.FindByAppID(applicationEntity.ID)
	if err != nil {
		return session.CurrentAppData{}, err
	}

	var appPlatforms []session.CurrentPlatformData
	for _, platformEntity := range platforms {
		appPlatforms = append(appPlatforms, session.CurrentPlatformData{
			PlatformID: platformEntity.ID,
			OS:         platformEntity.OS,
			AppSlug:    slug,
		})
	}

	// Desired platform order
	platformOrder := map[values.OS]int{
		values.Android: 1,
		values.IOS:     2,
		values.Windows: 3,
		values.MacOS:   4,
		values.Linux:   5,
	}

	// Sorting based on the platform order
	sort.Slice(appPlatforms, func(i, j int) bool {
		return platformOrder[appPlatforms[i].OS] < platformOrder[appPlatforms[j].OS]
	})

	return session.CurrentAppData{
		AppID:     applicationEntity.ID,
		AppName:   applicationEntity.Name,
		AppSlug:   applicationEntity.Slug,
		Platforms: appPlatforms,
	}, nil
}
