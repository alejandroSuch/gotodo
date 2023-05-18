package todo

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/application"
	"net/http"
)

type CompleteTodoController struct {
	completeTodo *application.CompleteTodo
}

func NewCompleteTodoController(
	completeTodo *application.CompleteTodo,
) *CompleteTodoController {
	return &CompleteTodoController{
		completeTodo: completeTodo,
	}
}

func (ctrl *CompleteTodoController) CompleteTodo(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	err := ctrl.completeTodo.Execute(application.CompleteTodoCommand{
		TodoId: c.Param("id"),
		UserId: userId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
