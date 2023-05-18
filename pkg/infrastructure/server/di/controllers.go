package di

import (
	"gotodo/pkg/rest/controllers/auth"
	"gotodo/pkg/rest/controllers/todo"
)

func InitializeLoginController() *auth.LoginController {
	login := InitializeLoginApplicationService()
	loginController := auth.NewLoginController(login, JwtSecretKey)
	return loginController
}

func InitializeRegisterController() *auth.RegisterController {
	registerUser := InitializeRegisterUserApplicationService()
	registerController := auth.NewRegisterController(registerUser)
	return registerController
}

func InitializeCreateTodoController() *todo.CreateTodoController {
	createTodo := InitializeCreateTodoApplicationService()
	createTodoController := todo.NewCreateTodoController(createTodo)
	return createTodoController
}

func InitializeCompleteTodoController() *todo.CompleteTodoController {
	completeTodo := InitializeCompleteTodoApplicationService()
	completeTodoController := todo.NewCompleteTodoController(completeTodo)
	return completeTodoController
}

func InitializeListTodosController() *todo.ListTodosController {
	userRepository := InitializeUserRepository()
	listTodosController := todo.NewListTodosController(userRepository)
	return listTodosController
}
