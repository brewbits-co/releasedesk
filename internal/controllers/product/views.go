package product

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/internal/views"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *appController) RenderDashboard(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	currentApp, err := c.service.GetCurrentAppData(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	overview, err := c.service.GetAppOverview(values.Slug(slug))
	if err != nil {
		// TODO: redirect to 404 page
		log.Println(err)
	}

	data := DashboardData{
		SessionData:         session.NewSessionData(r.Context()),
		CurrentAppData:      currentApp,
		SetupGuideCompleted: overview.SetupGuideCompleted,
	}

	tmpl, err := views.ParseTemplate(views.SidebarLayout, "templates/console/dashboard.gohtml")
	if err != nil {
		// TODO: redirect to error page
		log.Println(err)
	}

	_ = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	render.SetContentType(render.ContentTypeHTML)
}
