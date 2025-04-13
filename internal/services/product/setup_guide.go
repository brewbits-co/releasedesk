package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

func (s *service) ApplyProductSetupGuide(slug values.Slug, format values.VersionFormat, channels product.SetupChannelsOption, customChannels []string) error {
	productEntity, err := s.productRepo.FindBySlug(slug)
	if err != nil {
		return product.ErrProductNotFound
	}

	if productEntity.SetupGuideCompleted == true {
		return product.ErrSetupGuideAlreadyCompleted
	}

	var channelsToCreate []release.Channel
	if channels == product.ByMaturity {
		channelsToCreate = release.NewByMaturityChannels(productEntity.ID)
	}
	if channels == product.ByEnvironment {
		channelsToCreate = release.NewByEnvironmentChannels(productEntity.ID)
	}
	if channels == product.CustomChannels {
		channelsToCreate = make([]release.Channel, len(customChannels))
		for i := range customChannels {
			channelsToCreate[i] = release.NewChannel(productEntity.ID, customChannels[i], false)
		}
	}

	stepGuide := product.SetupGuide{
		ProductID:     productEntity.ID,
		VersionFormat: format,
		Channels:      channelsToCreate,
	}

	err = s.productRepo.SaveSetupGuide(stepGuide)
	if err != nil {
		return err
	}

	return nil
}
