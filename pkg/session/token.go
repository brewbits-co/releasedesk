package session

import (
	"github.com/go-chi/jwtauth/v5"
	"time"
)

// sessionDuration represents the duration for which a user session is considered valid, set to 72 hours.
const sessionDuration = time.Hour * 72

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

// CreateToken generates a JWT token for the given user with a validity period of 72 hours.
// The token contains user information such as user_id and username in its claims.
func CreateToken(userID int, username string) (string, error) {
	claims := make(map[string]interface{})
	claims["id"] = userID
	claims["username"] = username

	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, sessionDuration)

	_, tokenString, err := TokenAuth.Encode(claims)
	return tokenString, err
}
