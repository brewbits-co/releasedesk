package release

import (
	"github.com/brewbits-co/releasedesk/internal/services/product"
	"github.com/brewbits-co/releasedesk/internal/services/release"
	"net/http"
)

type ReleaseController interface {
	RenderReleaseList(w http.ResponseWriter, r *http.Request)
	RenderReleaseSummary(w http.ResponseWriter, r *http.Request)
	RenderReleaseNotes(w http.ResponseWriter, r *http.Request)
	HandleCreateRelease(w http.ResponseWriter, r *http.Request)
}

func NewReleaseController(service release.Service, productService product.Service) ReleaseController {
	return &releaseController{
		service:        service,
		productService: productService,
	}
}

type releaseController struct {
	service        release.Service
	productService product.Service
}
