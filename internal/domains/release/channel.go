package release

func NewChannel(appID int, name string, closed bool) Channel {
	return Channel{AppID: appID, Name: name, Closed: closed}
}

type Channel struct {
	// ID is the unique identifier of a Channel.
	ID int `db:"ID"`
	// AppID is the identifier of the app that this Channel belongs.
	AppID int `db:"AppID" json:"-"`
	// Name is a human-readable unique identifier of a Channel.
	Name string `db:"Name"`
	// Closed indicates whether the Channel is restricted to access only by invite, such as in a closed beta.
	Closed bool `db:"Closed" json:"-"`
}

func NewByMaturityChannels(appID int) []Channel {
	return []Channel{
		NewChannel(appID, "Canary", false),
		NewChannel(appID, "Beta", false),
		NewChannel(appID, "Stable", false),
	}
}

func NewByEnvironmentChannels(appID int) []Channel {
	return []Channel{
		NewChannel(appID, "Development", false),
		NewChannel(appID, "Staging", false),
		NewChannel(appID, "Production", false),
	}
}
