package product

import (
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"github.com/brewbits-co/releasedesk/internal/services/product"
	"net/http"
)

type AppController interface {
	HandleCreateApp(w http.ResponseWriter, r *http.Request)
	HandleAppSetupGuide(w http.ResponseWriter, r *http.Request)
	RenderDashboard(w http.ResponseWriter, r *http.Request)
}

// NewAppController creates a new instance of appController with the provided dependencies.
func NewAppController(service product.Service, platformService platform.Service) AppController {
	return &appController{
		service:         service,
		platformService: platformService,
	}
}

// appController implements the product.AppController.
type appController struct {
	service         product.Service
	platformService platform.Service
}
