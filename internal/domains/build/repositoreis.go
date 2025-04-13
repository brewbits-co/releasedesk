package build

type BuildRepository interface {
	Save(build *Build) error
	FindByAppID(appID int) ([]BasicInfo, error)
	GetByAppIDAndNumber(appID int, number int) (Build, error)
}
