package sql

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"xorm.io/xorm"
)

// NewBuildRepository is the constructor for buildRepository
func NewBuildRepository(engine *xorm.Engine) build.BuildRepository {
	return &buildRepository{engine: engine}
}

// buildRepository is the implementation of build.BuildRepository
type buildRepository struct {
	engine *xorm.Engine
}

func (r *buildRepository) Save(buildEntity *build.Build) error {
	session := r.engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	if _, err := session.Insert(buildEntity); err != nil {
		_ = session.Rollback()
		return err
	}

	for i := range buildEntity.Artifacts {
		buildEntity.Artifacts[i].BuildID = buildEntity.ID
		if _, err := session.Insert(&buildEntity.Artifacts[i]); err != nil {
			_ = session.Rollback()
			return err
		}
	}

	for key, value := range buildEntity.Metadata {
		metadata := build.BuildMetadata{
			BuildID: buildEntity.ID,
			Key:     key,
			Value:   value,
		}
		if _, err := session.Insert(&metadata); err != nil {
			_ = session.Rollback()
			return err
		}
	}

	if err := session.Commit(); err != nil {
		_ = session.Rollback()
		return err
	}

	return nil
}

func (r *buildRepository) FindByPlatformID(platformID int) ([]build.BasicInfo, error) {
	// Execute the database query
	var builds []build.BasicInfo
	err := r.engine.Desc("created_at").Table("build").Find(&builds, &build.BasicInfo{PlatformID: platformID})
	if err != nil {
		return nil, err // Return an error if the query fails
	}

	// Format timestamps for each build
	for i := range builds {
		builds[i].FormatAuditable()
	}

	return builds, nil
}

func (r *buildRepository) GetByPlatformIDAndNumber(platformID int, number int) (build.Build, error) {
	var buildDetails build.Build

	exist, err := r.engine.Where("platform_id = ? AND number = ?", platformID, number).Get(&buildDetails)
	if err != nil {
		return build.Build{}, err
	}
	if !exist {
		return build.Build{}, sql.ErrNoRows
	}

	buildDetails.Auditable.FormatAuditable()

	// Fetch build artifacts
	var artifacts []build.Artifact
	err = r.engine.Where("build_id = ?", buildDetails.ID).Find(&artifacts)
	if err != nil {
		return build.Build{}, err
	}
	buildDetails.Artifacts = artifacts

	// Fetch build metadata
	var metadataItems []build.BuildMetadata
	err = r.engine.Where("build_id = ?", buildDetails.ID).Find(&metadataItems)
	if err != nil {
		return build.Build{}, err
	}

	metadata := make(map[string]string)
	for _, item := range metadataItems {
		metadata[item.Key] = item.Value
	}
	buildDetails.Metadata = metadata

	return buildDetails, nil
}
