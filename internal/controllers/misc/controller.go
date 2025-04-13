package misc

import (
	"github.com/brewbits-co/releasedesk/internal/services/product"
	"net/http"
)

// MiscController defines the interface for
type MiscController interface {
	// RenderHomepage handles the rendering of the homepage by parsing the relevant HTML templates and writing them to the response.
	RenderHomepage(w http.ResponseWriter, r *http.Request)
}

// NewMiscController creates a new instance of authController with the provided dependencies.
func NewMiscController(service product.Service) MiscController {
	return &miscController{service: service}
}

// miscController implements the auth.MiscController.
type miscController struct {
	service product.Service
}
