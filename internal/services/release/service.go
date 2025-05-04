package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
)

// Service defines the interface for handling release-related use cases.
type Service interface {
	CreateRelease(slug values.Slug, info release.BasicInfo) (release.Release, error)
	ListReleasesByChannel(productID int, channelID int) ([]release.BasicInfo, error)
	GetReleaseSummary(productID int, version string) (release.Release, error)
	GetReleaseChannels(productID int) ([]release.Channel, error)
}

// NewReleaseService initializes a new instance of the release Service using the provided dependencies.
func NewReleaseService(releaseRepo release.ReleaseRepository, productRepo product.ProductRepository) Service {
	return &service{
		releaseRepo: releaseRepo,
		productRepo: productRepo,
	}
}

type service struct {
	releaseRepo release.ReleaseRepository
	productRepo product.ProductRepository
}
