package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/jmoiron/sqlx"
)

// NewBuildRepository is the constructor for buildRepository
func NewBuildRepository(db *sqlx.DB) build.BuildRepository {
	return &buildRepository{db: db}
}

// buildRepository is the implementation of build.BuildRepository
type buildRepository struct {
	db *sqlx.DB
}

func (r *buildRepository) Save(build *build.Build) error {
	_ = build.BeforeCreate()

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	q := `INSERT INTO Builds (AppID, Number, Version, CreatedAt, UpdatedAt) 
			VALUES (:AppID, :Number, :Version, :CreatedAt, :UpdatedAt)`

	result, err := tx.NamedExec(q, build)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	buildID, _ := result.LastInsertId()

	q = `INSERT INTO Artifacts (BuildID, Md5, Sha256, Sha512, Filename, MimeType, Size, Path, Architecture) 
			VALUES (:BuildID, :Md5, :Sha256, :Sha512, :Filename, :MimeType, :Size, :Path, :Architecture)`

	for _, artifact := range build.Artifacts {
		artifact.BuildID = int(buildID)

		result, err := tx.NamedExec(q, artifact)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		artifactID, _ := result.LastInsertId()
		artifact.ID = int(artifactID)
	}

	q = `INSERT INTO BuildMetadata (BuildID, Key, Value) VALUES ($1, $2, $3)`

	for key, value := range build.Metadata {
		_, err := tx.Exec(q, buildID, key, value)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if tx.Commit() != nil {
		_ = tx.Rollback()
		return err
	}

	_ = build.AfterCreate()
	return nil
}

func (r *buildRepository) FindByAppID(appID int) ([]build.BasicInfo, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT ID, AppID, Number, Version, CreatedAt, UpdatedAt FROM Builds WHERE AppID = $1 ORDER BY CreatedAt DESC", appID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the builds
	var builds []build.BasicInfo

	// Iterate over the result set
	for rows.Next() {
		var buildEntity build.BasicInfo
		// Map the row's data to the build struct
		if err := rows.StructScan(&buildEntity); err != nil {
			return nil, err // Return an error if mapping fails
		}
		buildEntity.FormatAuditable()
		builds = append(builds, buildEntity) // Add the build to the slice
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return builds, nil
}

func (r *buildRepository) GetByAppIDAndNumber(appID int, number int) (build.Build, error) {
	var buildDetails build.Build

	q := `SELECT * FROM Builds WHERE AppID = $1 AND Number = $2 LIMIT 1`

	err := r.db.QueryRowx(q, appID, number).StructScan(&buildDetails)
	if err != nil {
		return build.Build{}, err
	}

	buildDetails.Auditable.FormatAuditable()

	// Fetch build artifacts
	rows, err := r.db.Queryx("SELECT * FROM Artifacts WHERE BuildID = $1", buildDetails.ID)
	if err != nil {
		return build.Build{}, err
	}
	defer rows.Close()

	var artifacts []build.Artifact

	for rows.Next() {
		var artifactEntity build.Artifact
		if err := rows.StructScan(&artifactEntity); err != nil {
			return build.Build{}, err
		}
		artifacts = append(artifacts, artifactEntity)
	}

	if err := rows.Err(); err != nil {
		return build.Build{}, err
	}

	buildDetails.Artifacts = artifacts

	// Fetch build metadata
	rows, err = r.db.Queryx("SELECT Key, Value FROM BuildMetadata WHERE BuildID = $1", buildDetails.ID)
	if err != nil {
		return build.Build{}, err
	}
	defer rows.Close()

	metadata := make(map[string]string)

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return build.Build{}, err
		}

		metadata[key] = value
	}

	buildDetails.Metadata = metadata

	return buildDetails, nil
}
