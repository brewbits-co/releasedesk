package release

type ReleaseRepository interface {
	Save(release *Release) error
	FindByProductIDAndChannel(productID int, channelID int) ([]BasicInfo, error)
	FindChannelsByProductID(productID int) ([]Channel, error)
}
