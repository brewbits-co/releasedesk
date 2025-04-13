package product

import "github.com/brewbits-co/releasedesk/internal/values"

type ProductRepository interface {
	Save(product *Product) error
	Find() ([]Product, error)
	FindBySlug(slug values.Slug) (Product, error)
	Update(product Product) error
	Delete(product Product) error
	GetPlatformAvailability(product *Product) error
	SaveSetupGuide(guide SetupGuide) error
}
