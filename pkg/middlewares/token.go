package middlewares

import (
	"context"
	"github.com/brewbits-co/releasedesk/internal/services/auth"
	"net/http"
)

// APITokenAuthorization returns a middleware that handles the API token authorization.
func APITokenAuthorization(authSrv auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			_, token, _ := r.BasicAuth()

			if token != "90a514ab93e2c32fdd1072154b26a100" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			}

			sessionCtx := context.WithValue(r.Context(), "userID", 1)
			sessionCtx = context.WithValue(sessionCtx, "username", "admin")

			// Token is authenticated, pass it through
			next.ServeHTTP(w, r.WithContext(sessionCtx))
		}
		return http.HandlerFunc(hfn)
	}
}
