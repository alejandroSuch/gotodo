package rest

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/rest/controllers/auth"
	"gotodo/pkg/rest/controllers/todo"
)

func Start(
	loginCtrl *auth.LoginController,
	registerCtrl *auth.RegisterController,
	createTodoCtrl *todo.CreateTodoController,
	completeTodoCtrl *todo.CompleteTodoController,
	listTodosCtrl *todo.ListTodosController,
	authMiddleware gin.HandlerFunc,
) {

	router := gin.Default()

	router.POST("/login", loginCtrl.Login)
	router.POST("/register", registerCtrl.RegisterUser)
	router.POST("/todo", authMiddleware, createTodoCtrl.CreateTodo)
	router.GET("/todos", authMiddleware, listTodosCtrl.ListTodos)
	router.PATCH("/todos/:id/complete", authMiddleware, completeTodoCtrl.CompleteTodo)

	router.Run(":8080")
}
