package platform

import (
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/schemas"
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *platformController) HandleAddPlatform(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	slug := chi.URLParam(r, "slug")

	var platformRequest platform.BasicInfo

	err = utils.NewDecoder().Decode(&platformRequest, r.Form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	addedPlatform, err := c.service.AddPlatformToApp(values.Slug(slug), platformRequest)
	if err != nil {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, schemas.NewErrorResponse("A platform with the same identifier already exists.", []string{
			"Ensure only one platform is added per platform.",
			"Review existing apps to avoid duplicates before creating a new one.",
		}))
		return
	}

	log.Println(addedPlatform.ID)

	render.Status(r, http.StatusCreated)
}
