package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) ApplyAppSetupGuide(slug values.Slug, format values.VersionFormat, channels app.SetupChannelsOption, customChannels []string) error {
	appEntity, err := s.appRepo.FindBySlug(slug)
	if err != nil {
		return app.ErrAppNotFound
	}

	if appEntity.SetupGuideCompleted == true {
		return app.ErrSetupGuideAlreadyCompleted
	}

	var channelsToCreate []release.Channel
	if channels == app.ByMaturity {
		channelsToCreate = release.NewByMaturityChannels(appEntity.ID)
	}
	if channels == app.ByEnvironment {
		channelsToCreate = release.NewByEnvironmentChannels(appEntity.ID)
	}
	if channels == app.CustomChannels {
		channelsToCreate = make([]release.Channel, len(customChannels))
		for i := range customChannels {
			channelsToCreate[i] = release.NewChannel(appEntity.ID, customChannels[i], false)
		}
	}

	stepGuide := app.SetupGuide{
		AppID:         appEntity.ID,
		VersionFormat: format,
		Channels:      channelsToCreate,
	}

	err = s.appRepo.SaveSetupGuide(stepGuide)
	if err != nil {
		return err
	}

	return nil
}
