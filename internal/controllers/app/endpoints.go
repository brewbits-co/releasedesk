package app

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/schemas"
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *appController) HandleCreateApp(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	slug := chi.URLParam(r, "slug")

	var appCreationRequest app.BasicInfo

	err = utils.NewDecoder().Decode(&appCreationRequest, r.Form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	newApp, err := c.service.CreateApp(values.Slug(slug), appCreationRequest)
	if err != nil {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, schemas.NewErrorResponse("A app with the same identifier already exists.", []string{
			"Ensure only one app is added per platform.",
			"Review existing apps to avoid duplicates before creating a new one.",
		}))
		return
	}

	log.Println(newApp.ID)

	render.Status(r, http.StatusCreated)
}
