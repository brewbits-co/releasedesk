package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/schemas"
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

// HandleUpdateBasicInfo handles the update of a release's BasicInfo
// Only the TargetChannel and Status fields can be updated
func (c *releaseController) HandleUpdateBasicInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	slug := chi.URLParam(r, "slug")
	version := chi.URLParam(r, "version")

	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, schemas.NewErrorResponse("App not found", []string{err.Error()}))
		return
	}

	var updateRequest release.BasicInfo

	err = utils.NewDecoder().Decode(&updateRequest, r.Form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	_, err = c.service.UpdateReleaseBasicInfo(currentApp.AppID, version, updateRequest)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, schemas.NewErrorResponse("Failed to update release", []string{err.Error()}))
		return
	}

	// Redirect back to the release summary page
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
