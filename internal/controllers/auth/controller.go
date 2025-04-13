package auth

import (
	"github.com/brewbits-co/releasedesk/internal/services/auth"
	"net/http"
)

// AuthController defines the interface for managing user authentication, including login and logout operations and rendering appropriate views.
type AuthController interface {
	// HandleLogin processes the login request by validating user credentials and generating a session token if authentication is successful.
	HandleLogin(w http.ResponseWriter, r *http.Request)
	// HandleLogout invalidates the user session by clearing the session cookie and redirects to the login page.
	HandleLogout(w http.ResponseWriter, r *http.Request)
	// RenderLogin handles the rendering of the login page by parsing the relevant HTML templates and writing them to the response.
	RenderLogin(w http.ResponseWriter, r *http.Request)
}

// NewAuthController creates a new instance of authController with the provided dependencies.
func NewAuthController(service auth.Service) AuthController {
	return &authController{service: service}
}

// authController implements the auth.AuthController.
type authController struct {
	service auth.Service
}
