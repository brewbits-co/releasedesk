package app

import "github.com/brewbits-co/releasedesk/internal/values"

type AppRepository interface {
	Save(app *App) error
	FindByProductID(productID int) ([]App, error)
	GetByProductSlugAndPlatform(slug values.Slug, platform values.Platform) (App, error)
}
