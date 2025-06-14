package release

func NewChannel(appID int, name string, closed bool) Channel {
	return Channel{AppID: appID, Name: name, Closed: closed}
}

type Channel struct {
	// ID is the unique identifier of a Channel.
	ID int `xorm:"'ID' pk autoincr"`
	// AppID is the identifier of the app that this Channel belongs.
	AppID int `json:"-"`
	// Name is a human-readable unique identifier of a Channel.
	Name string `xorm:"not null"`
	// Closed indicates whether the Channel is restricted to access only by invite, such as in a closed beta.
	Closed bool `xorm:"not null default false" json:"-"`
}

// NewByMaturityChannels creates a predefined set of channels based on software maturity levels.
// It creates three channels: Canary (for early testing), Beta (for broader testing),
// and Stable (for production-ready releases), all associated with the specified appID.
func NewByMaturityChannels(appID int) []Channel {
	return []Channel{
		NewChannel(appID, "Canary", false),
		NewChannel(appID, "Beta", false),
		NewChannel(appID, "Stable", false),
	}
}

// NewByEnvironmentChannels creates a predefined set of channels based on deployment environments.
// It creates three channels: Development (for dev environment), Staging (for pre-production),
// and Production (for live environment), all associated with the specified appID.
func NewByEnvironmentChannels(appID int) []Channel {
	return []Channel{
		NewChannel(appID, "Development", false),
		NewChannel(appID, "Staging", false),
		NewChannel(appID, "Production", false),
	}
}
