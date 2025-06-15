package sql

import (
	"errors"
	"github.com/brewbits-co/releasedesk/internal/domains/user"
	"xorm.io/xorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// NewUserRepository is the constructor for userRepository
func NewUserRepository(engine *xorm.Engine) user.UserRepository {
	return &userRepository{engine: engine}
}

// userRepository is the implementation of user.UserRepository
type userRepository struct {
	engine *xorm.Engine
}

func (r *userRepository) GetByID(id int) (user.User, error) {
	var userEntity user.User
	found, err := r.engine.ID(id).Get(&userEntity)
	if err != nil || !found {
		return user.User{}, ErrUserNotFound
	}
	userEntity.FormatAuditable()

	return userEntity, nil
}

func (r *userRepository) GetByUsername(username string) (user.User, error) {
	var userEntity user.User
	found, err := r.engine.Where("username = ?", username).Get(&userEntity)
	if err != nil || !found {
		return user.User{}, ErrUserNotFound
	}
	userEntity.FormatAuditable()

	return userEntity, nil
}
