package application

import "gotodo/pkg/domain"

type CompleteTodo struct {
	userRepo     domain.UserRepository
	nextIdentity domain.NextIdentity
}

type CompleteTodoCommand struct {
	UserId string
	TodoId string
}

func NewCompleteTodoApplicationService(
	userRepo domain.UserRepository,
	nextIdentity domain.NextIdentity,
) *CompleteTodo {
	return &CompleteTodo{
		userRepo:     userRepo,
		nextIdentity: nextIdentity,
	}
}

func (s CompleteTodo) Execute(c CompleteTodoCommand) error {
	user, err := s.userRepo.FindById(c.UserId)
	if err != nil {
		return err
	}

	user.CompleteTodo(c.TodoId)

	err = s.userRepo.Save(user, s.nextIdentity)
	if err != nil {
		return err
	}

	return nil
}
