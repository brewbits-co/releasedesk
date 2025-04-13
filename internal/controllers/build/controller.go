package build

import (
	"github.com/brewbits-co/releasedesk/internal/services/build"
	"github.com/brewbits-co/releasedesk/internal/services/product"
	"net/http"
)

// BuildController defines the interface for managing build operations and rendering appropriate views.
type BuildController interface {
	HandleBuildUpload(w http.ResponseWriter, r *http.Request)
	HandleArtifactDownload(w http.ResponseWriter, r *http.Request)
	RenderBuildList(w http.ResponseWriter, r *http.Request)
	RenderBuildDetails(w http.ResponseWriter, r *http.Request)
	RenderBuildMetadata(w http.ResponseWriter, r *http.Request)
}

func NewBuildController(service build.Service, productService product.Service) BuildController {
	return &buildController{
		service:        service,
		productService: productService,
	}
}

type buildController struct {
	service        build.Service
	productService product.Service
}
