package app

import "github.com/brewbits-co/releasedesk/internal/values"

type PlatformRepository interface {
	Save(platform *Platform) error
	FindByAppID(appID int) ([]Platform, error)
	GetByAppSlugAndOS(slug values.Slug, os values.OS) (Platform, error)
}
