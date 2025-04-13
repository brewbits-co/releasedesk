package auth

import (
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/render"
	"net/http"
)

func (c *authController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := c.service.Login(username, password)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	tokenString, err := session.CreateToken(user.ID, user.Username)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	cookie := session.CreateLoginCookie(tokenString)

	http.SetCookie(w, &cookie)
	render.Status(r, http.StatusOK)
}

func (c *authController) HandleLogout(w http.ResponseWriter, r *http.Request) {
	cookie := session.CreateLogoutCookie()
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
