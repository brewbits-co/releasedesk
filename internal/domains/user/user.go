package user

import (
	"database/sql"
	"errors"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

// NewUser will create a new User entity and return an error if the email is not valid or the password is not strong
// The password parameter must be a plain-text password, it will be converted to a bcrypt hashed password.
func NewUser(username string, firstName string, lastName string, email string, password string) (User, error) {
	if !validator.IsEmail(email) {
		return User{}, ErrInvalidEmail
	}

	if !validator.IsPasswordStrong(password) {
		return User{}, errors.New("weak password")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return User{
		Auditable: fields.NewAuditable(),
		Username:  username,
		Email:     sql.NullString{String: email, Valid: true},
		Password:  values.HashedPassword(hash),
		FirstName: sql.NullString{String: firstName, Valid: true},
		LastName:  sql.NullString{String: lastName, Valid: true},
		Role:      values.Viewer,
	}, nil
}

type User struct {
	validator.BaseValidator `xorm:"-"`
	fields.Auditable        `xorm:"extends"`
	// ID is the unique identifier of a User.
	ID int `xorm:"pk autoincr"`
	// Username is a unique identifier of the user.
	Username string `xorm:"not null unique"`
	// Email is the email address of the user.
	Email sql.NullString
	// Password is the bcrypt hashed password of the user.
	// It is excluded from JSON serialization for security reasons.
	Password values.HashedPassword `json:"-" xorm:"not null"`
	// FirstName is the first name of the user.
	FirstName sql.NullString
	// LastName is the last name of the user.
	LastName sql.NullString
	// Role represents the user's role in the system.
	Role values.Role `xorm:"not null"`
}

// IsValid checks if the current user information follows the pre-defined business rules
func (u *User) IsValid() error {
	if validator.IsEmail(u.Email.String) {
		return ErrInvalidEmail
	}
	if u.Role == 0 {
		return ErrMustHaveRole
	}
	if validator.IsAnyEmpty(u.Username, u.FirstName.String, u.LastName.String, string(u.Password)) {
		return ErrEmptyField
	}
	return nil
}
