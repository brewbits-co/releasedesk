package release

import (
	"github.com/brewbits-co/releasedesk/internal/services/app"
	"github.com/brewbits-co/releasedesk/internal/services/release"
	"net/http"
)

type ReleaseController interface {
	RenderReleaseList(w http.ResponseWriter, r *http.Request)
	RenderReleaseSummary(w http.ResponseWriter, r *http.Request)
	RenderReleaseNotes(w http.ResponseWriter, r *http.Request)
	HandleCreateRelease(w http.ResponseWriter, r *http.Request)
	HandleSaveReleaseNotes(w http.ResponseWriter, r *http.Request)
	HandleUpdateBasicInfo(w http.ResponseWriter, r *http.Request)
}

func NewReleaseController(service release.Service, appService app.Service) ReleaseController {
	return &releaseController{
		service:    service,
		appService: appService,
	}
}

type releaseController struct {
	service    release.Service
	appService app.Service
}
