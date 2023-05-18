package application

import "gotodo/pkg/domain"

type CreateTodo struct {
	userRepo     domain.UserRepository
	nextIdentity domain.NextIdentity
}

type CreateTodoCommand struct {
	UserId      string
	Description string
}

func NewCreateTodoApplicationService(
	userRepo domain.UserRepository,
	nextIdentity domain.NextIdentity,
) *CreateTodo {
	return &CreateTodo{
		userRepo:     userRepo,
		nextIdentity: nextIdentity,
	}
}

func (s CreateTodo) Execute(c CreateTodoCommand) error {
	user, err := s.userRepo.FindById(c.UserId)
	if err != nil {
		return err
	}

	user.AddTodo(c.Description)

	err = s.userRepo.Save(user, s.nextIdentity)
	if err != nil {
		return err
	}

	return nil
}
