package auth

import (
	"github.com/brewbits-co/releasedesk/internal/views"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *authController) RenderLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := views.ParseTemplate(views.NoLayout, "templates/console/login.gohtml")
	if err != nil {
		log.Println(err)
	}

	_ = tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	render.SetContentType(render.ContentTypeHTML)
}
