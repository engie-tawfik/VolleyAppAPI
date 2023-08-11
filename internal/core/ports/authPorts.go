package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Login(email string) domain.LoginTeam
}

type AuthService interface {
	Login(email, password string) domain.AuthResponse
	CreateTokens(teamId string) domain.AuthResponse
}

type AuthHandler interface {
	Login(c *gin.Context)
	RefreshTokens(c *gin.Context)
	RegisterAuthRoutes(rg *gin.RouterGroup)
}

type AuthMiddleware interface {
	RequireAuth(c *gin.Context)
	RequireRefresh(c *gin.Context)
}
