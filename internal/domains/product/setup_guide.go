package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

type SetupGuide struct {
	AppID         int
	VersionFormat values.VersionFormat
	Channels      []release.Channel
}

type SetupChannelsOption string

const (
	// CustomChannels is tailored to your product's unique requirements.
	CustomChannels SetupChannelsOption = "CustomChannels"
	// ByMaturity represent the stability of the release and the intended audience.
	ByMaturity SetupChannelsOption = "ByMaturity"
	// ByEnvironment  map releases to environments for effective deployment control.
	ByEnvironment SetupChannelsOption = "ByEnvironment"
)
