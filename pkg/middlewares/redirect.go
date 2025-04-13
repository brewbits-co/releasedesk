package middlewares

import (
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http"
	"strings"
)

// RedirectOnUnauthorized returns a middleware that redirects unauthorized requests to the login page or returns a 401 error.
func RedirectOnUnauthorized(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			options := ja.ValidateOptions()

			if err != nil || token == nil || jwt.Validate(token, options...) != nil {
				log.Println(err)
				handleUnauthorized(w, r)
				return
			}

			err = session.RefreshCookieToken(w, token)
			if err != nil {
				log.Println(err)
				handleUnauthorized(w, r)
			}

			sessionCtx, err := session.CreateSessionContext(r.Context(), token)
			if err != nil {
				log.Println(err)
				handleUnauthorized(w, r)
			}

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r.WithContext(sessionCtx))
		}
		return http.HandlerFunc(hfn)
	}
}

func handleUnauthorized(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/internal") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	} else {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}
}
