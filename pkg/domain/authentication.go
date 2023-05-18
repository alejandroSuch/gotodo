package domain

type Authentication struct {
	ID           string
	Username     string
	PasswordHash string
}

func NewAuthentication(id string, username string, passwordHash string) Authentication {
	return Authentication{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
	}
}
