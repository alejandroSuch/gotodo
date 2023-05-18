package domain

type PasswordMatches func(password string, hash string) bool
