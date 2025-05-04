package platform

import (
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"net/http"
)

type PlatformController interface {
	HandleAddPlatform(w http.ResponseWriter, r *http.Request)
}

func NewPlatformController(service platform.Service) PlatformController {
	return &platformController{service: service}
}

type platformController struct {
	service platform.Service
}
