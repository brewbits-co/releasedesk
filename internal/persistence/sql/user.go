package sql

import (
	"errors"
	"github.com/brewbits-co/releasedesk/internal/domains/user"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// NewUserRepository is the constructor for userRepository
func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepository{db: db}
}

// userRepository is the implementation of user.UserRepository
type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) FindByID(id int) (user.User, error) {
	var userEntity user.User
	err := r.db.QueryRowx("SELECT * FROM users WHERE ID = $1 LIMIT 1", id).StructScan(&userEntity)
	if err != nil {
		return user.User{}, ErrUserNotFound
	}

	return userEntity, nil
}

func (r *userRepository) FindByUsername(username string) (user.User, error) {
	var userEntity user.User
	err := r.db.QueryRowx("SELECT * FROM users WHERE Username = $1 LIMIT 1", username).StructScan(&userEntity)
	if err != nil {
		return user.User{}, ErrUserNotFound
	}

	return userEntity, nil
}
