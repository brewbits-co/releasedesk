package sql

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"xorm.io/xorm"
)

// NewReleaseRepository is the constructor for releaseRepository
func NewReleaseRepository(engine *xorm.Engine) release.ReleaseRepository {
	return &releaseRepository{engine: engine}
}

// releaseRepository is the implementation of release.ReleaseRepository
type releaseRepository struct {
	engine *xorm.Engine
}

func (r *releaseRepository) FindChannelsByAppID(appID int) ([]release.Channel, error) {
	var channels []release.Channel
	err := r.engine.Where("app_id = ?", appID).Asc("id").Find(&channels)
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *releaseRepository) Save(release *release.Release) error {
	_, err := r.engine.Insert(release)
	if err != nil {
		return err
	}

	return nil
}

func (r *releaseRepository) Update(release *release.Release) error {
	_, err := r.engine.Update(release)
	if err != nil {
		return err
	}

	return nil
}

func (r *releaseRepository) FindByAppIDAndChannel(appID int, channelID int) ([]release.BasicInfo, error) {
	// Execute the database query using xorm
	var releases []release.BasicInfo
	err := r.engine.Table("release").Where("app_id = ? AND target_channel = ?", appID, channelID).
		Desc("created_at").
		Find(&releases)
	if err != nil {
		return nil, err
	}

	// Format the auditable fields for each release
	for i := range releases {
		releases[i].FormatAuditable()
	}

	return releases, nil
}

func (r *releaseRepository) GetByAppIDAndVersion(appID int, version string) (release.Release, error) {
	var releaseSummary release.Release

	// Execute the database query using xorm
	has, err := r.engine.Where("app_id = ? AND version = ?", appID, version).Get(&releaseSummary)
	if err != nil {
		return release.Release{}, err
	}
	if !has {
		return release.Release{}, sql.ErrNoRows
	}

	releaseSummary.Auditable.FormatAuditable()

	return releaseSummary, nil
}

// GetByID retrieves a release by its ID
func (r *releaseRepository) GetByID(id int) (release.Release, error) {
	var releaseSummary release.Release

	// Execute the database query using xorm
	has, err := r.engine.ID(id).Get(&releaseSummary)
	if err != nil {
		return release.Release{}, err
	}
	if !has {
		return release.Release{}, sql.ErrNoRows
	}

	releaseSummary.Auditable.FormatAuditable()

	return releaseSummary, nil
}
