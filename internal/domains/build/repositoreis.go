package build

type BuildRepository interface {
	Save(build *Build) error
	FindByPlatformID(platformID int) ([]BasicInfo, error)
	GetByPlatformIDAndNumber(platformID int, number int) (Build, error)
}
