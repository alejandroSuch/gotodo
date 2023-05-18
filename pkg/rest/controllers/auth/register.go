package auth

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/application"
	"net/http"
)

type RegisterController struct {
	registerUser *application.RegisterUser
}

func NewRegisterController(registerUser *application.RegisterUser) *RegisterController {
	return &RegisterController{
		registerUser: registerUser,
	}
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *RegisterController) RegisterUser(ctx *gin.Context) {
	var input RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.registerUser.Execute(application.RegisterUserCommand{
		Username: input.Username,
		Name:     input.Name,
		Password: input.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
