package controllers

import (
	"fmt"
	"log"
	"net/http"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
	"volleyapp/infrastructure/errors"

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
	models.RegisterUserValidators()
	return &AuthController{
		gin:               gin,
		authService:       authService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleWare,
	}
}

func (a *AuthController) InitAuthRoutes() {
	authBasePath := fmt.Sprintf("%s/auth", config.BasePath)
	authRoute := a.gin.Group(authBasePath, a.headersMiddleware.RequireApiKey)
	authRoute.POST("/users/create", a.CreateUser)
	authRoute.POST("/login", a.Login)
	authRoute.POST(
		"/refresh", a.authMiddleware.RequireRefresh, a.RefreshTokens,
	)
}

func (a *AuthController) Login(c *gin.Context) {
	var authData models.Auth

	if err := c.ShouldBindJSON(&authData); err != nil {
		log.Println("auth controller - error in Login:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Auth controller - Login request - email:", authData.Email)

	authResponse, err := a.authService.Login(authData.Email, authData.Password)
	if err != nil {
		log.Println("Auth controller - error in Login:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Access",
		authResponse.AccessToken,
		config.JwtExpireMins*60,
		"",
		"",
		true,
		true,
	)
	c.SetCookie(
		"Refresh",
		authResponse.Refreshtoken,
		config.JwtExpireMins*60*2,
		"",
		"",
		true,
		true,
	)
	log.Println("Auth controller - tokens were set in cookies")
	c.Status(http.StatusOK)
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	userId, _ := c.Get("userId")
	log.Println("Auth controller - RefreshTokens request - userId:", userId)
	authResponse, err := a.authService.CreateTokens(int(userId.(float64)))
	if err != nil {
		log.Println("Auth controller - error in RefreshTokens:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Access",
		authResponse.AccessToken,
		config.JwtExpireMins*60,
		"",
		"",
		true,
		true,
	)
	c.SetCookie(
		"Refresh",
		authResponse.Refreshtoken,
		config.JwtExpireMins*60*2,
		"",
		"",
		true,
		true,
	)
	log.Println("Auth controller - tokens were set in cookies")
	c.Status(http.StatusOK)
}

func (a *AuthController) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println("Auth controller - unable to process user:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Auth controller - CreateUser request - email:", newUser.Email)

	userId, err := a.authService.CreateUser(newUser)
	if err != nil {
		log.Println("Auth controller - error in CreateUser:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Auth controller - user created with id:", userId)
	response := models.Response{
		Message: "User was successfully created",
		Data:    map[string]int{"userId": userId},
	}
	c.JSON(http.StatusCreated, response)
}
