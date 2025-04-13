package values

import (
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword string

// CompareWith checks if the provided plainTextPassword matches the bcrypt-hashed password stored in the User struct.
func (p HashedPassword) CompareWith(plainTextPassword string) bool {
	hashErr := bcrypt.CompareHashAndPassword([]byte(p), []byte(plainTextPassword))
	return hashErr == nil
}
