package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/jmoiron/sqlx"
)

// NewAppRepository is the constructor for appRepository
func NewAppRepository(db *sqlx.DB) app.PlatformRepository {
	return &appRepository{db: db}
}

// appRepository is the implementation of app.PlatformRepository
type appRepository struct {
	db *sqlx.DB
}

func (r *appRepository) Save(app *app.Platform) error {
	_ = app.BeforeCreate()

	q := `INSERT INTO Apps (ProductID, Name, OS, CreatedAt, UpdatedAt) 
			VALUES (:ProductID, :Name, :OS, :CreatedAt, :UpdatedAt)`

	exec, err := r.db.NamedExec(q, app)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	app.ID = int(insertId)

	_ = app.AfterCreate()
	return nil
}

func (r *appRepository) FindByAppID(productID int) ([]app.Platform, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT * FROM Apps WHERE ProductID = $1", productID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the apps
	var apps []app.Platform

	// Iterate over the result set
	for rows.Next() {
		var appEntity app.Platform
		// Map the row's data to the app struct
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

func (r *appRepository) GetByAppSlugAndOS(slug values.Slug, platform values.OS) (app.Platform, error) {
	var appInfo app.Platform

	q := `SELECT Apps.* FROM Apps
	JOIN Products ON Apps.ProductID = Products.ID
	WHERE Products.Slug = $1 AND Apps.OS = $2
	LIMIT 1`

	err := r.db.QueryRowx(q, slug, platform).StructScan(&appInfo)
	if err != nil {
		return app.Platform{}, err
	}

	return appInfo, nil
}
