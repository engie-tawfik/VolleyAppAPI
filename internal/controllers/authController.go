package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var badRequestResponse = &domain.Response{
	ErrorCode: http.StatusBadRequest,
	Message:   "Bad request",
	Data:      nil,
}

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
	response := domain.Response{
		Message: "",
		Data:    nil,
	}
	if err := c.ShouldBindJSON(&authData); err != nil {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Error in Login: %s", err.Error()),
		)
		c.AbortWithStatusJSON(badRequestResponse.ErrorCode, badRequestResponse)
		return
	}
	logger.Logger.Info("Login request for " + authData.Email)
	logger.Logger.Debug(fmt.Sprintf("Login data: %+v", authData))
	authResponse := a.authService.Login(authData.Email, authData.Password)
	if authResponse.AccessToken == "" || authResponse.Refreshtoken == "" {
		logger.Logger.Error(
			fmt.Sprintf("[AUTH CONTROLLER] Error in Login: tokens were not created"),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
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
	return
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	userId, _ := c.Get("userId")
	logger.Logger.Info(
		fmt.Sprintf(
			"[AUTH CONTROLLER] RefreshTokens request userId: %d", userId,
		),
	)
	authResponse := a.authService.CreateTokens(userId.(int))
	if authResponse.AccessToken == "" || authResponse.Refreshtoken == "" {
		logger.Logger.Error(
			fmt.Sprintf(
				"[AUTH CONTROLLER] Error in RefreshTokens: tokens were not created",
			),
		)
		c.AbortWithStatusJSON(badRequestResponse.ErrorCode, badRequestResponse)
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
			fmt.Sprintf("Unable to process User: %s", err.Error()),
		)
		var ve validator.ValidationErrors
		// Check for validation errors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				errorMsg := domain.GetTeamErrorMsg(fe)
				if errorMsg != "" {
					logger.Logger.Error(
						fmt.Sprintf("Validation error for User: %s", errorMsg),
					)
					break
				}
			}
		}
		response.Message = ("Bad request")
		c.AbortWithStatusJSON(badRequestResponse.ErrorCode, badRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf("Request for CreateUser. User: %s", newUser.Email),
	)
	logger.Logger.Debug(
		fmt.Sprintf("New User data: %s, %s", newUser.Email, newUser.Password),
	)
	userId, err := a.authService.CreateUser(newUser)
	if err != nil {
		logger.Logger.Error("Error in create user: " + err.Error())
		c.AbortWithStatusJSON(badRequestResponse.ErrorCode, badRequestResponse)
		return
	}
	logger.Logger.Debug(
		fmt.Sprintf("[AUTH CONTROLLER] User created - New userId: ", userId),
	)
	response.Message = "User created"
	response.Data = userId
	c.JSON(http.StatusCreated, response)
}
