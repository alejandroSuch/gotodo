package di

import (
	"gotodo/pkg/application"
)

func InitializeLoginApplicationService() *application.Login {
	authenticationRepository := InitializeAuthRepository()
	userRepository := InitializeUserRepository()
	passwordMatches := InitializePasswordMatches()
	login := application.NewLoginApplicationService(authenticationRepository, userRepository, passwordMatches)
	return login
}

func InitializeRegisterUserApplicationService() *application.RegisterUser {
	authenticationRepository := InitializeAuthRepository()
	userRepository := InitializeUserRepository()
	passwordHasher := InitializePasswordHasher()
	nextIdentity := InitializeNextIdentity()
	registerUser := application.NewRegisterUserApplicationService(authenticationRepository, userRepository, passwordHasher, nextIdentity)
	return registerUser
}

func InitializeCreateTodoApplicationService() *application.CreateTodo {
	userRepository := InitializeUserRepository()
	nextIdentity := InitializeNextIdentity()
	createTodo := application.NewCreateTodoApplicationService(userRepository, nextIdentity)
	return createTodo
}

func InitializeCompleteTodoApplicationService() *application.CompleteTodo {
	userRepository := InitializeUserRepository()
	nextIdentity := InitializeNextIdentity()
	completeTodo := application.NewCompleteTodoApplicationService(userRepository, nextIdentity)
	return completeTodo
}
