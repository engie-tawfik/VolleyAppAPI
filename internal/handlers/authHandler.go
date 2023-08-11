package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"volleyapp/config"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService    ports.AuthService
	authMiddleware ports.AuthMiddleware
}

var _ ports.AuthHandler = (*AuthHandler)(nil)

func NewAuthHandler(authService ports.AuthService, authMiddleware ports.AuthMiddleware) *AuthHandler {
	return &AuthHandler{
		authService:    authService,
		authMiddleware: authMiddleware,
	}
}

func (a *AuthHandler) Login(c *gin.Context) {
	var authData domain.Auth
	response := domain.Response{
		Message: "",
		Data:    nil,
	}
	if err := c.ShouldBindJSON(&authData); err != nil {
		config.Logger.Info("Wrong login credentials")
		response.Message = "Wrong login credentials."
		c.JSON(http.StatusBadRequest, response)
		return
	}
	config.Logger.Info("Login request for " + authData.Email)
	config.Logger.Debug(fmt.Sprintf("Login data: %+v", authData))
	authResponse := a.authService.Login(authData.Email, authData.Password)
	if authResponse.AccessToken == "" || authResponse.Refreshtoken == "" {
		response.Message = "Wrong login credentials."
		c.JSON(http.StatusBadRequest, response)
		return
	}
	tokenLife := os.Getenv("JWT_TOKEN_EXPIRE_MINUTES")
	tokenLifeInt, _ := strconv.Atoi(tokenLife)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Access",
		authResponse.AccessToken,
		tokenLifeInt*60,
		"",
		"",
		true,
		true,
	)
	c.SetCookie(
		"Refresh",
		authResponse.Refreshtoken,
		tokenLifeInt*60*2,
		"",
		"",
		true,
		true,
	)
	// authResponse := domain.AuthResponse{
	// 	AccessToken:  accessToken,
	// 	Refreshtoken: refreshToken,
	// }
	c.Status(http.StatusOK)
}

func (a *AuthHandler) RefreshTokens(c *gin.Context) {
	teamId, _ := c.Get("teamId")
	config.Logger.Info("Request for RefreshTokens. TeamId: " + teamId.(string))
	authResponse := a.authService.CreateTokens(teamId.(string))
	if authResponse.AccessToken == "" || authResponse.Refreshtoken == "" {
		response := domain.Response{
			Message: "Wrong login credentials.",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	tokenLife := os.Getenv("JWT_TOKEN_EXPIRE_MINUTES")
	tokenLifeInt, _ := strconv.Atoi(tokenLife)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Access",
		authResponse.AccessToken,
		tokenLifeInt*60,
		"",
		"",
		true,
		true,
	)
	c.SetCookie(
		"Refresh",
		authResponse.Refreshtoken,
		tokenLifeInt*60*2,
		"",
		"",
		true,
		true,
	)
	// authResponse := domain.AuthResponse{
	// 	AccessToken:  accessToken,
	// 	Refreshtoken: refreshToken,
	// }
	c.Status(http.StatusOK)
}

func (a *AuthHandler) RegisterAuthRoutes(rg *gin.RouterGroup) {
	authRoute := rg.Group("/auth")
	authRoute.POST("/login", a.Login)
	authRoute.POST("/refresh", a.authMiddleware.RequireRefresh, a.RefreshTokens)
}
