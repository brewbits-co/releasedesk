package release

func NewChannel(productID int, name string, closed bool) Channel {
	return Channel{ProductID: productID, Name: name, Closed: closed}
}

type Channel struct {
	// ID is the unique identifier of a Channel.
	ID int `db:"ID"`
	// ProductID is the identifier of the product that this Channel belongs.
	ProductID int `db:"PlatformID" json:"-"`
	// Name is a human-readable unique identifier of a Channel.
	Name string `db:"Name"`
	// Closed indicates whether the Channel is restricted to access only by invite, such as in a closed beta.
	Closed bool `db:"Closed" json:"-"`
}

func NewByMaturityChannels(productID int) []Channel {
	return []Channel{
		NewChannel(productID, "Canary", false),
		NewChannel(productID, "Beta", false),
		NewChannel(productID, "Stable", false),
	}
}

func NewByEnvironmentChannels(productID int) []Channel {
	return []Channel{
		NewChannel(productID, "Development", false),
		NewChannel(productID, "Staging", false),
		NewChannel(productID, "Production", false),
	}
}
