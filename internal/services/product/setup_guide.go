package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) ApplyAppSetupGuide(slug values.Slug, format values.VersionFormat, channels product.SetupChannelsOption, customChannels []string) error {
	appEntity, err := s.appRepo.FindBySlug(slug)
	if err != nil {
		return product.ErrAppNotFound
	}

	if appEntity.SetupGuideCompleted == true {
		return product.ErrSetupGuideAlreadyCompleted
	}

	var channelsToCreate []release.Channel
	if channels == product.ByMaturity {
		channelsToCreate = release.NewByMaturityChannels(appEntity.ID)
	}
	if channels == product.ByEnvironment {
		channelsToCreate = release.NewByEnvironmentChannels(appEntity.ID)
	}
	if channels == product.CustomChannels {
		channelsToCreate = make([]release.Channel, len(customChannels))
		for i := range customChannels {
			channelsToCreate[i] = release.NewChannel(appEntity.ID, customChannels[i], false)
		}
	}

	stepGuide := product.SetupGuide{
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
