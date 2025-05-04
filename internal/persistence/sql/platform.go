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

func (r *platformRepository) Save(app *platform.Platform) error {
	_ = app.BeforeCreate()

	q := `INSERT INTO Platforms (PlatformID, OS, CreatedAt, UpdatedAt) 
			VALUES (:PlatformID, :OS, :CreatedAt, :UpdatedAt)`

	exec, err := r.db.NamedExec(q, app)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	app.ID = int(insertId)

	_ = app.AfterCreate()
	return nil
}

func (r *platformRepository) FindByAppID(productID int) ([]platform.Platform, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT * FROM Platforms WHERE PlatformID = $1", productID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the apps
	var apps []platform.Platform

	// Iterate over the result set
	for rows.Next() {
		var appEntity platform.Platform
		// Map the row's data to the platform struct
		if err := rows.StructScan(&appEntity); err != nil {
			return nil, err // Return an error if mapping fails
		}
		appEntity.FormatAuditable()
		apps = append(apps, appEntity) // Add the product to the slice
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apps, nil // Return the list of apps
}

func (r *platformRepository) GetByAppSlugAndOS(slug values.Slug, os values.OS) (platform.Platform, error) {
	var platformInfo platform.Platform

	q := `SELECT Platforms.* FROM Platforms
	JOIN Products ON Platforms.PlatformID = Products.ID
	WHERE Products.Slug = $1 AND Platforms.OS = $2
	LIMIT 1`

	err := r.db.QueryRowx(q, slug, os).StructScan(&platformInfo)
	if err != nil {
		return platform.Platform{}, err
	}

	return platformInfo, nil
}
