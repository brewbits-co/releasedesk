package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/jmoiron/sqlx"
)

// NewPlatformRepository is the constructor for platformRepository
func NewPlatformRepository(db *sqlx.DB) platform.PlatformRepository {
	return &platformRepository{db: db}
}

// platformRepository is the implementation of platform.PlatformRepository
type platformRepository struct {
	db *sqlx.DB
}

func (r *platformRepository) Save(platform *platform.Platform) error {
	_ = platform.BeforeCreate()

	q := `INSERT INTO Platforms (AppID, OS, CreatedAt, UpdatedAt) 
			VALUES (:AppID, :OS, :CreatedAt, :UpdatedAt)`

	exec, err := r.db.NamedExec(q, platform)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	platform.ID = int(insertId)

	_ = platform.AfterCreate()
	return nil
}

func (r *platformRepository) FindByAppID(appID int) ([]platform.Platform, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT * FROM Platforms WHERE AppID = $1", appID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the platforms
	var platforms []platform.Platform

	// Iterate over the result set
	for rows.Next() {
		var platformEntity platform.Platform
		// Map the row's data to the platform struct
		if err := rows.StructScan(&platformEntity); err != nil {
			return nil, err // Return an error if mapping fails
		}
		platformEntity.FormatAuditable()
		platforms = append(platforms, platformEntity) // Add the app to the slice
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return platforms, nil // Return the list of platforms
}

func (r *platformRepository) GetByAppSlugAndOS(slug values.Slug, os values.OS) (platform.Platform, error) {
	var platformInfo platform.Platform

	q := `SELECT Platforms.* FROM Platforms
	JOIN Apps ON Platforms.AppID = Apps.ID
	WHERE Apps.Slug = $1 AND Platforms.OS = $2
	LIMIT 1`

	err := r.db.QueryRowx(q, slug, os).StructScan(&platformInfo)
	if err != nil {
		return platform.Platform{}, err
	}

	return platformInfo, nil
}
