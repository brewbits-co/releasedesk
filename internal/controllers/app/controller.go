package app

import (
	"github.com/brewbits-co/releasedesk/internal/services/app"
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"net/http"
)

type AppController interface {
	HandleCreateApp(w http.ResponseWriter, r *http.Request)
	HandleAppSetupGuide(w http.ResponseWriter, r *http.Request)
	RenderDashboard(w http.ResponseWriter, r *http.Request)
}

// NewAppController creates a new instance of appController with the provided dependencies.
func NewAppController(service app.Service, platformService platform.Service) AppController {
	return &appController{
		service:         service,
		platformService: platformService,
	}
}

// appController implements the app.AppController.
type appController struct {
	service         app.Service
	platformService platform.Service
}
