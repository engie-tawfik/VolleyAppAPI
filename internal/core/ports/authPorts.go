package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	GetUserByEmail(email string) (domain.User, error)
	SaveNewUser(user domain.User) (int, error)
}

type AuthService interface {
	Login(email, password string) (domain.AuthResponse, error)
	CreateTokens(userId int) (domain.AuthResponse, error)
	CreateUser(user domain.User) (int, error)
}

type AuthController interface {
	Login(c *gin.Context)
	CreateUser(c *gin.Context)
	RefreshTokens(c *gin.Context)
	InitAuthRoutes()
}

type AuthMiddleware interface {
	RequireAuth(c *gin.Context)
	RequireRefresh(c *gin.Context)
}
