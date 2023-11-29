package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	GetUserByEmail(string) (domain.User, error)
	SaveNewUser(domain.User) (int, error)
}

type AuthService interface {
	Login(string, string) (domain.AuthResponse, error)
	CreateTokens(int) (domain.AuthResponse, error)
	CreateUser(domain.User) (int, error)
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
