package build

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/internal/views"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"strconv"
)

func (c *buildController) RenderBuildList(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentPlatform := session.NewCurrentPlatformData(currentApp, values.OS(platform))

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/build_list.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	builds, err := c.service.GetPlatformBuilds(currentPlatform.PlatformID)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildListData{
		CurrentPage:     "Builds",
		SessionData:     session.NewSessionData(r.Context()),
		CurrentAppData:  currentApp,
		CurrentPlatform: currentPlatform,
		Builds:          builds,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *buildController) RenderBuildDetails(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentPlatform := session.NewCurrentPlatformData(currentApp, values.OS(platform))

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/build_details.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	buildDetails, err := c.service.GetBuildDetails(currentPlatform.PlatformID, number)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildDetailsData{
		CurrentPage:     "Builds",
		SessionData:     session.NewSessionData(r.Context()),
		CurrentAppData:  currentApp,
		CurrentPlatform: currentPlatform,
		Build:           buildDetails,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *buildController) RenderBuildMetadata(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentPlatform := session.NewCurrentPlatformData(currentApp, values.OS(platform))

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/build_metadata.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	buildDetails, err := c.service.GetBuildDetails(currentPlatform.PlatformID, number)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildDetailsData{
		CurrentPage:     "Builds",
		SessionData:     session.NewSessionData(r.Context()),
		CurrentAppData:  currentApp,
		CurrentPlatform: currentPlatform,
		Build:           buildDetails,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}
