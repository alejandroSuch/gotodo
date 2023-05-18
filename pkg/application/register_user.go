package application

import "gotodo/pkg/domain"

type RegisterUser struct {
	hash     domain.PasswordHasher
	nextId   domain.NextIdentity
	authRepo domain.AuthenticationRepository
	userRepo domain.UserRepository
}

type RegisterUserCommand struct {
	Name     string
	Username string
	Password string
}

func NewRegisterUserApplicationService(
	authRepo domain.AuthenticationRepository,
	userRepo domain.UserRepository,
	hash domain.PasswordHasher,
	nextId domain.NextIdentity,
) *RegisterUser {
	return &RegisterUser{
		authRepo: authRepo,
		userRepo: userRepo,
		hash:     hash,
		nextId:   nextId,
	}
}

func (r RegisterUser) Execute(c RegisterUserCommand) (*domain.User, error) {
	hash, err := r.hash(c.Password)
	if err != nil {
		return nil, err
	}

	auth := domain.NewAuthentication(r.nextId(), c.Username, hash)
	err = r.authRepo.Save(&auth)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(r.nextId(), auth.ID, c.Name)
	err = r.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
