package domain

type PasswordHasher func(string) (string, error)
