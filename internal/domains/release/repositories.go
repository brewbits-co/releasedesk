package release

type ReleaseRepository interface {
	Save(release *Release) error
	FindByAppIDAndChannel(appID int, channelID int) ([]BasicInfo, error)
	FindChannelsByAppID(appID int) ([]Channel, error)
	GetByAppIDAndVersion(appID int, version string) (Release, error)
}

type ReleaseNotesRepository interface {
	Save(releaseNotes *ReleaseNotes) error
	FindByReleaseID(releaseID int) (ReleaseNotes, error)
	FindChangelogsByReleaseID(releaseID int) ([]Changelog, error)
}
