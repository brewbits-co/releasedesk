package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
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
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	channels, err := c.service.GetReleaseChannels(currentApp.AppID)
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

	releases, err := c.service.ListReleasesByChannel(currentApp.AppID, currentChannelID)
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := ReleaseListData{
		CurrentPage:      "Releases",
		SessionData:      session.NewSessionData(r.Context()),
		CurrentAppData:   currentApp,
		Releases:         releases,
		Channels:         channels,
		CurrentChannelID: currentChannelID,
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
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/release_summary.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	channels, err := c.service.GetReleaseChannels(currentApp.AppID)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	version := chi.URLParam(r, "version")
	releaseSummary, err := c.service.GetReleaseSummary(currentApp.AppID, version)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	data := ReleaseSummaryData{
		CurrentPage:    "Releases",
		SessionData:    session.NewSessionData(r.Context()),
		CurrentAppData: currentApp,
		Channels:       channels,
		Release:        releaseSummary,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}

func (c *releaseController) RenderReleaseNotes(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentApp, err := c.appService.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	version := chi.URLParam(r, "version")
	releaseSummary, err := c.service.GetReleaseSummary(currentApp.AppID, version)
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	releaseNotes, err := c.service.GetReleaseNotes(releaseSummary.ID)
	if err != nil {
		// This might be a new release without notes, so we'll create an empty one
		releaseNotes = release.NewReleaseNotes(releaseSummary.ID, "")
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/release_notes.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	data := ReleaseNotesData{
		CurrentPage:    "Releases",
		SessionData:    session.NewSessionData(r.Context()),
		CurrentAppData: currentApp,
		Release:        releaseSummary,
		ReleaseNotes:   releaseNotes,
	}

	err = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
	render.SetContentType(render.ContentTypeHTML)
}
