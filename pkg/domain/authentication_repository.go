package domain

import "errors"

var (
	ErrAuthNotFound = errors.New("authentication not found")
	ErrUnauthorized = errors.New("unauthorized")
)

type AuthenticationRepository interface {
	Save(auth *Authentication) error
	GetByUsername(username string) (*Authentication, error)
	GetByUsernameAndPassword(username string, password string, passwordComparer PasswordMatches) (*Authentication, error)
}
