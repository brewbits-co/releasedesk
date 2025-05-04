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
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentApp := session.NewCurrentAppData(currentProduct, values.OS(platform))

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/build_list.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	builds, err := c.service.GetPlatformBuilds(currentApp.AppID)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildListData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
		CurrentApp:         currentApp,
		Builds:             builds,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *buildController) RenderBuildDetails(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentApp := session.NewCurrentAppData(currentProduct, values.OS(platform))

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

	buildDetails, err := c.service.GetBuildDetails(currentApp.AppID, number)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildDetailsData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
		CurrentApp:         currentApp,
		Build:              buildDetails,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *buildController) RenderBuildMetadata(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	platform := chi.URLParam(r, "platform")
	currentApp := session.NewCurrentAppData(currentProduct, values.OS(platform))

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

	buildDetails, err := c.service.GetBuildDetails(currentApp.AppID, number)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := BuildDetailsData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
		CurrentApp:         currentApp,
		Build:              buildDetails,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}
