package application

import "gotodo/pkg/domain"

type Login struct {
	authRepo        domain.AuthenticationRepository
	userRepo        domain.UserRepository
	passwordMatches domain.PasswordMatches
}

type LoginCommand struct {
	Username string
	Password string
}

func NewLoginApplicationService(
	authRepo domain.AuthenticationRepository,
	userRepo domain.UserRepository,
	passwordMatches domain.PasswordMatches,
) *Login {
	return &Login{
		authRepo:        authRepo,
		userRepo:        userRepo,
		passwordMatches: passwordMatches,
	}
}

func (l Login) Execute(c LoginCommand) (*domain.User, error) {
	auth, err := l.authRepo.GetByUsernameAndPassword(c.Username, c.Password, l.passwordMatches)
	if err != nil {
		return nil, err
	}

	user, err := l.userRepo.FindByAuthenticationID(auth.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
