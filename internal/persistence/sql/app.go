package sql

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/values"
	"xorm.io/xorm"
)

// NewApplicationRepository is the constructor for appRepository
func NewApplicationRepository(engine *xorm.Engine) app.AppRepository {
	return &appRepository{engine: engine}
}

// appRepository is the implementation of app.AppRepository
type appRepository struct {
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

func (r *appRepository) GetBySlug(slug values.Slug) (app.App, error) {
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
	session := r.engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	_, err := session.Table("app").
		Where("id = ?", guide.AppID).
		Update(map[string]interface{}{
			"version_format":        guide.VersionFormat,
			"setup_guide_completed": true,
		})
	if err != nil {
		session.Rollback()
		return err
	}

	for _, channel := range guide.Channels {
		_, err := session.Insert(&channel)
		if err != nil {
			session.Rollback()
			return err
		}
	}

	return session.Commit()
}

func (r *appRepository) GetPlatformAvailability(app *app.App) error {
	platforms := make([]string, 0)
	err := r.engine.Table("platform").
		Where("app_id = ?", app.ID).
		Cols("os").
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
