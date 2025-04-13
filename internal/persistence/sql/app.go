package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/jmoiron/sqlx"
)

// NewAppRepository is the constructor for appRepository
func NewAppRepository(db *sqlx.DB) app.AppRepository {
	return &appRepository{db: db}
}

// appRepository is the implementation of app.AppRepository
type appRepository struct {
	db *sqlx.DB
}

func (r *appRepository) Save(app *app.App) error {
	_ = app.BeforeCreate()

	q := `INSERT INTO Apps (ProductID, Name, Platform, CreatedAt, UpdatedAt) 
			VALUES (:ProductID, :Name, :Platform, :CreatedAt, :UpdatedAt)`

	exec, err := r.db.NamedExec(q, app)
	if err != nil {
		return err
	}

	insertId, _ := exec.LastInsertId()
	app.ID = int(insertId)

	_ = app.AfterCreate()
	return nil
}

func (r *appRepository) FindByProductID(productID int) ([]app.App, error) {
	// Execute the database query
	rows, err := r.db.Queryx("SELECT * FROM Apps WHERE ProductID = $1", productID)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the cursor is closed when the function exits

	// Declare a slice to store the apps
	var apps []app.App

	// Iterate over the result set
	for rows.Next() {
		var appEntity app.App
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

func (r *appRepository) GetByProductSlugAndPlatform(slug values.Slug, platform values.Platform) (app.App, error) {
	var appInfo app.App

	q := `SELECT Apps.* FROM Apps
	JOIN Products ON Apps.ProductID = Products.ID
	WHERE Products.Slug = $1 AND Apps.Platform = $2
	LIMIT 1`

	err := r.db.QueryRowx(q, slug, platform).StructScan(&appInfo)
	if err != nil {
		return app.App{}, err
	}

	return appInfo, nil
}
