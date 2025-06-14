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
	_, err := r.engine.Insert(app)
	if err != nil {
		return err
	}

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
	_, err := r.engine.ID(app.ID).Update(&app)
	if err != nil {
		return err
	}

	return nil
}

func (r *appRepository) Delete(app app.App) error {
	_, err := r.engine.ID(app.ID).Delete(&app)
	if err != nil {
		return err
	}
	return nil
}

func (r *appRepository) SaveSetupGuide(guide app.SetupGuide) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	q := "UPDATE app SET version_format = $1, setup_guide_completed = true WHERE ID = $2"
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
	platforms := make([]string, 0)
	err := r.engine.Table("Platforms").
		Where("AppID = ?", app.ID).
		Cols("OS").
		Find(&platforms)
	if err != nil {
		return err
	}

	for _, platform := range platforms {
		switch platform {
		case "Android":
			app.HasAndroid = true
		case "iOS":
			app.HasIOS = true
		case "Windows":
			app.HasWindows = true
		case "Linux":
			app.HasLinux = true
		case "macOS":
			app.HasMacOS = true
		}
	}

	return nil
}
