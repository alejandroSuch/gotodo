package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Create(user *User) error
	Save(user *User, nextIdentity NextIdentity) error
	FindById(id string) (*User, error)
	FindByAuthenticationID(authId string) (*User, error)
}
