package user

type UserRepository interface {
	FindByID(id int) (User, error)
	FindByUsername(username string) (User, error)
}
