package sql

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/values"
	"xorm.io/xorm"
)

// NewPlatformRepository is the constructor for platformRepository
func NewPlatformRepository(engine *xorm.Engine) platform.PlatformRepository {
	return &platformRepository{engine: engine}
}

// platformRepository is the implementation of platform.PlatformRepository
type platformRepository struct {
	engine *xorm.Engine
}

func (r *platformRepository) Save(platform *platform.Platform) error {
	_, err := r.engine.Insert(platform)
	if err != nil {
		return err
	}
	return nil
}

func (r *platformRepository) FindByAppID(appID int) ([]platform.Platform, error) {
	var platforms []platform.Platform
	err := r.engine.Where("app_id = ?", appID).Find(&platforms)
	if err != nil {
		return nil, err
	}

	for i := range platforms {
		platforms[i].FormatAuditable()
	}

	return platforms, nil
}

func (r *platformRepository) GetByAppSlugAndOS(slug values.Slug, os values.OS) (platform.Platform, error) {
	var platformInfo platform.Platform
	exist, err := r.engine.Join("INNER", "app", "platform.app_id = app.id").
		Where("app.slug = ? AND platform.os = ?", slug, os).
		Get(&platformInfo)
	if err != nil {
		return platform.Platform{}, err
	}
	if !exist {
		return platform.Platform{}, sql.ErrNoRows
	}

	platformInfo.FormatAuditable()
	return platformInfo, nil
}
