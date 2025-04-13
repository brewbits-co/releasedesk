package app

import (
	"github.com/brewbits-co/releasedesk/internal/services/app"
	"net/http"
)

type AppController interface {
	HandleCreateApp(w http.ResponseWriter, r *http.Request)
}

func NewAppController(service app.Service) AppController {
	return &appController{service: service}
}

type appController struct {
	service app.Service
}
