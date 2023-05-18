package todo

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/application"
	"net/http"
)

type CreateTodoController struct {
	createTodo *application.CreateTodo
}

func NewCreateTodoController(
	createTodo *application.CreateTodo,
) *CreateTodoController {
	return &CreateTodoController{
		createTodo: createTodo,
	}
}

type CreateTodoInput struct {
	Description string `json:"description" binding:"required"`
}

func (ctrl *CreateTodoController) CreateTodo(c *gin.Context) {
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err := ctrl.createTodo.Execute(application.CreateTodoCommand{
		UserId:      userId,
		Description: input.Description,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
