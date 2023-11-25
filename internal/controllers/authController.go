package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/errors"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	gin               *gin.Engine
	authService       ports.AuthService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.AuthController = (*AuthController)(nil)

func NewAuthController(
	gin *gin.Engine,
	authService ports.AuthService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleWare ports.HeadersMiddleware,
) *AuthController {
	domain.RegisterUserValidators()
	return &AuthController{
		gin:               gin,
		authService:       authService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleWare,
	}
}

func (a *AuthController) InitAuthRoutes() {
	authBasePath := fmt.Sprintf("%s/auth", os.Getenv("BASE_PATH"))
	authRoute := a.gin.Group(authBasePath, a.headersMiddleware.RequireApiKey)
	authRoute.POST("/users/create", a.CreateUser)
	authRoute.POST("/login", a.Login)
	authRoute.POST(
		"/refresh", a.authMiddleware.RequireRefresh, a.RefreshTokens,
	)
}

func (a *AuthController) Login(c *gin.Context) {
	var authData domain.Auth

	if err := c.ShouldBindJSON(&authData); err != nil {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Error in Login: %s", err),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(fmt.Sprintf("[AUTH CONTROLLER] Login request: %s", authData.Email))
	logger.Logger.Debug(fmt.Sprintf("[AUTH CONTROLLER] Login data: %s", authData))
	authResponse, err := a.authService.Login(authData.Email, authData.Password)
	if err != nil {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Error in Login: %s", err),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
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

	c.Status(http.StatusOK)
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	userId, _ := c.Get("userId")
	logger.Logger.Info(
		fmt.Sprintf("[AUTH CONTROLLER] RefreshTokens request userId: %v", userId),
	)
	authResponse := a.authService.CreateTokens(int(userId.(float64)))
	if authResponse.AccessToken == "" || authResponse.Refreshtoken == "" {
		logger.Logger.Error(
			"[AUTH CONTROLLER] Error in RefreshTokens: tokens were not created",
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
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
	c.Status(http.StatusOK)
}

func (a *AuthController) CreateUser(c *gin.Context) {
	var newUser domain.User
	response := domain.Response{
		Message: "",
		Data:    nil,
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Unable to process User: %s", err),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf("[AUTH CONTROLLER] Request for CreateUser. User: %s", newUser.Email),
	)
	logger.Logger.Debug(
		fmt.Sprintf("[AUTH CONTROLLER] New User data: %s, %s", newUser.Email, newUser.Password),
	)
	userId, err := a.authService.CreateUser(newUser)
	if err != nil {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Error in create user: %s", err),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Debug(
		fmt.Sprintf("[AUTH CONTROLLER] User created - New userId: %d", userId),
	)
	response.Message = "User created"
	response.Data = userId
	c.JSON(http.StatusCreated, response)
}
