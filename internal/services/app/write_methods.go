package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
)

func (s *service) CreateApp(info app.BasicInfo) (app.App, error) {
	info.SetupGuideCompleted = false
	info.Slug.Format()

	newApp := app.NewApp(info)

	err := newApp.IsValid()
	if err != nil {
		return app.App{}, err
	}

	err = s.appRepo.Save(&newApp)
	if err != nil {
		return app.App{}, err
	}

	return newApp, nil
}
