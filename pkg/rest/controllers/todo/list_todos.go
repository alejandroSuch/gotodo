package todo

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/domain"
	"net/http"
)

type ListTodosController struct {
	userRepo domain.UserRepository
}

func NewListTodosController(userRepo domain.UserRepository) *ListTodosController {
	return &ListTodosController{
		userRepo: userRepo,
	}
}

type Todo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (ctrl ListTodosController) ListTodos(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := ctrl.userRepo.FindById(userId)
	if err == domain.ErrUserNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos := make([]Todo, len(user.Todos))
	for i, t := range user.Todos {
		todos[i] = Todo{
			Id:          t.ID,
			Description: t.Description,
			Completed:   t.Status == domain.TodoStatusCompleted,
		}
	}

	c.JSON(http.StatusOK, todos)
}
