package session

import (
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"
)

// sessionCookieName is the name of the HTTP cookie used to store the session's JWT token.
const sessionCookieName = "jwt"

// CreateLoginCookie generates an HTTP cookie with the specified JWT token string, used to maintain user session.
func CreateLoginCookie(tokenString string) http.Cookie {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(sessionDuration),
		HttpOnly: true,
	}
	return cookie
}

// CreateLogoutCookie creates an HTTP cookie to invalidate the session by setting an expired JWT token.
func CreateLogoutCookie() http.Cookie {
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		MaxAge:   -1,
	}
	return cookie
}

func RefreshCookieToken(w http.ResponseWriter, token jwt.Token) error {
	id, username, err := ExtractTokenInformation(token)
	if err != nil {
		return errors.New("failed to extract token information")
	}

	tokenString, err := CreateToken(id, username)
	if err != nil {
		return errors.New("failed to create token")
	}

	cookie := CreateLoginCookie(tokenString)
	http.SetCookie(w, &cookie)
	return nil
}

func ExtractTokenInformation(token jwt.Token) (int, string, error) {
	id, ok := token.Get("id")
	if !ok {
		return 0, "", errors.New("id claim not found")
	}

	username, ok := token.Get("username")
	if !ok {
		return 0, "", errors.New("username claim not found")
	}
	return int(id.(float64)), username.(string), nil
}
