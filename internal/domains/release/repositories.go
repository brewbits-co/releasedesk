package release

type ReleaseRepository interface {
	Save(release *Release) error
	FindByAppIDAndChannel(appID int, channelID int) ([]BasicInfo, error)
	FindChannelsByAppID(appID int) ([]Channel, error)
	GetByAppIDAndVersion(appID int, version string) (Release, error)
}
