package di

import (
	"gotodo/pkg/domain"
	"gotodo/pkg/domain/impl"
)

func InitializePasswordMatches() domain.PasswordMatches {
	return impl.BcryptPasswordMatches
}

func InitializePasswordHasher() domain.PasswordHasher {
	return impl.BcryptPasswordHasher
}

func InitializeNextIdentity() domain.NextIdentity {
	return impl.UUIDNextIdentity
}
