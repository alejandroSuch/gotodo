package auth

import (
	"gotodo/pkg/application"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	login     *application.Login
	secretKey string
}

func NewLoginController(login *application.Login, secretKey string) *LoginController {
	return &LoginController{
		login:     login,
		secretKey: secretKey,
	}
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ctrl *LoginController) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.login.Execute(application.LoginCommand{
		Username: input.Username,
		Password: input.Password,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("my-super-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
