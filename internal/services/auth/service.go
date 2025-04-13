package auth

import (
	"github.com/brewbits-co/releasedesk/internal/domains/user"
)

// Service defines the interface for handling authentication-related use cases.
type Service interface {
	// Login authenticates a user based on the provided username and password.
	// It returns the authenticated user's account information or any potential error.
	Login(username string, password string) (user.User, error)
}

// NewAuthService initializes a new instance of the authentication Service using the provided dependencies.
func NewAuthService(userRepo user.UserRepository) Service {
	srv := &service{
		userRepo: userRepo,
	}

	return srv
}

// service implements the auth.Service.
type service struct {
	userRepo user.UserRepository
}
