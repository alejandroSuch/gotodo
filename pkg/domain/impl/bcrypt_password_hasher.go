package impl

import "golang.org/x/crypto/bcrypt"

func BcryptPasswordHasher(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	return string(h), err
}
