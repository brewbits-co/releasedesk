package user

type UserRepository interface {
	GetByID(id int) (User, error)
	GetByUsername(username string) (User, error)
}
