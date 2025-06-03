package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/jmoiron/sqlx"
	"xorm.io/xorm"
)

// NewApplicationRepository is the constructor for appRepository
func NewApplicationRepository(db *sqlx.DB, engine *xorm.Engine) app.AppRepository {
	return &appRepository{db: db, engine: engine}
}

// appRepository is the implementation of app.AppRepository
type appRepository struct {
	db     *sqlx.DB
	engine *xorm.Engine
}

func (r *appRepository) Save(app *app.App) error {
	_ = app.BeforeCreate()

	_, err := r.engine.Insert(app)
	if err != nil {
		return err
	}

	_ = app.AfterCreate()
	return nil
}

func (r *appRepository) Find() ([]app.App, error) {
	var apps []app.App
	err := r.engine.Find(&apps)
	if err != nil {
		return nil, err
	}

	for _, row := range apps {
		row.FormatAuditable()
	}

	return apps, nil // Return the list of apps
}

func (r *appRepository) FindBySlug(slug values.Slug) (app.App, error) {
	var p app.App
	p.Slug = slug
	_, err := r.engine.Get(&p)
	if err != nil {
		return app.App{}, err
	}

	return p, err
}

func (r *appRepository) Update(app app.App) error {
	_ = app.BeforeUpdate()

	q := `UPDATE Apps SET 
			Name = :Name, 
			Slug = :Slug, 
			Description = :Description,
			Private = :Private,
           	VersionFormat = :VersionFormat, 
			SetupGuideCompleted = :SetupGuideCompleted WHERE ID = :ID`

	_, err := r.db.NamedExec(q, app)
	if err != nil {
		return err
	}

	_ = app.AfterUpdate()
	return nil
}

func (r *appRepository) Delete(app app.App) error {
	//TODO implement me
	panic("implement me")
}

func (r *appRepository) SaveSetupGuide(guide app.SetupGuide) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	q := "UPDATE Apps SET VersionFormat = $1, SetupGuideCompleted = true WHERE ID = $2"
	_, err = tx.Exec(q, guide.VersionFormat, guide.AppID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	q = "INSERT INTO Channels (Name, AppID, Closed) VALUES (:Name, :AppID, :Closed)"
	for _, channel := range guide.Channels {
		_, err := tx.NamedExec(q, channel)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if tx.Commit() != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (r *appRepository) GetPlatformAvailability(app *app.App) error {
	q := `SELECT 
    EXISTS (
        SELECT 1 
        FROM Platforms a 
        WHERE a.AppID = p.ID AND a.OS = 'Android'
    ) AS HasAndroid,
    EXISTS (
        SELECT 1 
        FROM Platforms a 
        WHERE a.AppID = p.ID AND a.OS = 'iOS'
    ) AS HasIOS,
    EXISTS (
        SELECT 1 
        FROM Platforms a 
        WHERE a.AppID = p.ID AND a.OS = 'Windows'
    ) AS HasWindows,
    EXISTS (
        SELECT 1 
        FROM Platforms a 
        WHERE a.AppID = p.ID AND a.OS = 'Linux'
    ) AS HasLinux,
    EXISTS (
        SELECT 1 
        FROM Platforms a 
        WHERE a.AppID = p.ID AND a.OS = 'macOS'
    ) AS HasMacOS FROM Apps p WHERE ID = $1`

	row := r.db.QueryRow(q, app.ID)

	err := row.Scan(
		&app.HasAndroid,
		&app.HasIOS,
		&app.HasWindows,
		&app.HasLinux,
		&app.HasMacOS,
	)
	if err != nil {
		return err
	}

	return nil
}
