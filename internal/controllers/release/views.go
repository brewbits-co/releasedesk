package release

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

func (c *releaseController) RenderReleaseList(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	channels, err := c.service.GetReleaseChannels(currentProduct.ProductID)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	channel := r.URL.Query().Get("channel")

	var currentChannelID int
	if channel == "" {
		currentChannelID = channels[0].ID
	} else {
		currentChannelID, _ = strconv.Atoi(channel)
	}

	releases, err := c.service.ListReleasesByChannel(currentProduct.ProductID, currentChannelID)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := ReleaseListData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
		Releases:           releases,
		Channels:           channels,
		CurrentChannelID:   currentChannelID,
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/release_list.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)

}

func (c *releaseController) RenderReleaseSummary(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/release_summary.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	channels, err := c.service.GetReleaseChannels(currentProduct.ProductID)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	version := chi.URLParam(r, "version")
	releaseSummary, err := c.service.GetReleaseSummary(currentProduct.ProductID, version)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	data := ReleaseSummaryData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
		Channels:           channels,
		Release:            releaseSummary,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *releaseController) RenderReleaseNotes(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentProduct, err := c.productService.GetCurrentProductData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/release_notes.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := ReleaseNotesData{
		SessionData:        session.NewSessionData(r.Context()),
		CurrentProductData: currentProduct,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}
