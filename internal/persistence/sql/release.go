package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/jmoiron/sqlx"
)

// NewReleaseRepository is the constructor for releaseRepository
func NewReleaseRepository(db *sqlx.DB) release.ReleaseRepository {
	return &releaseRepository{db: db}
}

// releaseRepository is the implementation of release.ReleaseRepository
type releaseRepository struct {
	db *sqlx.DB
}

func (r *releaseRepository) FindChannelsByAppID(appID int) ([]release.Channel, error) {
	rows, err := r.db.Queryx(`SELECT ID, Name, AppID, Closed FROM Channels WHERE AppID = $1 ORDER BY ID`, appID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []release.Channel

	for rows.Next() {
		var channelEntity release.Channel
		if err := rows.StructScan(&channelEntity); err != nil {
			return nil, err
		}
		channels = append(channels, channelEntity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

func (r *releaseRepository) Save(release *release.Release) error {
	_ = release.BeforeCreate()

	q := `INSERT INTO Releases (AppID, Version, TargetChannel, Status, CreatedAt, UpdatedAt) 
			VALUES (:AppID, :Version, :TargetChannel, :Status, :CreatedAt, :UpdatedAt)`

	exec, err := r.db.NamedExec(q, release)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	release.ID = int(insertId)

	_ = release.AfterCreate()
	return nil
}

func (r *releaseRepository) FindByAppIDAndChannel(appID int, channelID int) ([]release.BasicInfo, error) {
	// Execute the database query
	q := `SELECT ID, AppID, Version, TargetChannel, Status, CreatedAt, UpdatedAt 
			FROM Releases WHERE AppID = $1 AND TargetChannel = $2 ORDER BY CreatedAt DESC`
	rows, err := r.db.Queryx(q, appID, channelID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the releases
	var releases []release.BasicInfo

	// Iterate over the result set
	for rows.Next() {
		var releaseEntity release.BasicInfo
		// Map the row's data to the build struct
		if err := rows.StructScan(&releaseEntity); err != nil {
			return nil, err // Return an error if mapping fails
		}
		releaseEntity.FormatAuditable()
		releases = append(releases, releaseEntity) // Add the build to the slice
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return releases, nil
}

func (r *releaseRepository) GetByAppIDAndVersion(appID int, version string) (release.Release, error) {
	var releaseSummary release.Release

	// Execute the database query
	q := `SELECT ID, AppID, Version, TargetChannel, Status, CreatedAt, UpdatedAt 
			FROM Releases WHERE AppID = $1 AND Version = $2 LIMIT 1`

	err := r.db.QueryRowx(q, appID, version).StructScan(&releaseSummary)
	if err != nil {
		return release.Release{}, err
	}

	releaseSummary.Auditable.FormatAuditable()

	return releaseSummary, nil
}
