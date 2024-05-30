package ports

import (
	"volleyapp/domain/models"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	GetUserByEmail(string) (models.User, error)
	SaveNewUser(models.User) (int, error)
}

type AuthService interface {
	Login(string, string) (models.AuthResponse, error)
	CreateTokens(int) (models.AuthResponse, error)
	CreateUser(models.User) (int, error)
}

type AuthController interface {
	Login(*gin.Context)
	CreateUser(*gin.Context)
	RefreshTokens(*gin.Context)
	InitAuthRoutes()
}

type AuthMiddleware interface {
	RequireAuth(*gin.Context)
	RequireRefresh(*gin.Context)
}
