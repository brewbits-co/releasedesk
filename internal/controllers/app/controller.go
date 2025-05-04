package app

import (
	"github.com/brewbits-co/releasedesk/internal/services/app"
	"net/http"
)

type PlatformController interface {
	HandleAddPlatform(w http.ResponseWriter, r *http.Request)
}

func NewPlatformController(service app.Service) PlatformController {
	return &platformController{service: service}
}

type platformController struct {
	service app.Service
}
