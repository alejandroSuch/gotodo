package di

import (
	"github.com/gin-gonic/gin"
	"gotodo/pkg/rest/middleware"
)

func InitializeAuthMiddleware() gin.HandlerFunc {
	handlerFunc := middleware.AuthMiddleware(JwtSecretKey)
	return handlerFunc
}
