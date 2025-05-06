package misc

import (
	"github.com/brewbits-co/releasedesk/internal/views"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *miscController) RenderHomepage(w http.ResponseWriter, r *http.Request) {
	apps, err := c.service.GetUserAccessibleApps(r.Context().Value("userID").(int))
	data := HomepageData{
		SessionData: session.NewSessionData(r.Context()),
		Apps:        apps,
	}

	tmpl, err := views.ParseTemplate(views.NavbarLayout, "templates/console/homepage.gohtml")
	if err != nil {
		log.Println(err)
	}

	_ = tmpl.ExecuteTemplate(w, "index.gohtml", data)
	render.SetContentType(render.ContentTypeHTML)
}
