package impl

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptPasswordMatches(password string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}
