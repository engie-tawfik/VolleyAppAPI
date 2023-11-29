package controllers

import (
	"fmt"
	"net/http"
	"os"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/errors"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gin               *gin.Engine
	gameService       ports.GameService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.GameController = (*GameController)(nil)

func NewGameController(
	gin *gin.Engine,
	gameService ports.GameService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleware ports.HeadersMiddleware,
) *GameController {
	return &GameController{
		gin:               gin,
		gameService:       gameService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (g *GameController) InitGameRoutes() {
	gameBasePath := fmt.Sprintf("%s/games", os.Getenv("BASE_PATH"))
	gameRoute := g.gin.Group(
		gameBasePath,
		g.headersMiddleware.RequireApiKey,
		g.authMiddleware.RequireAuth,
	)
	gameRoute.POST("/create", g.CreateGame)
}

func (g *GameController) CreateGame(c *gin.Context) {
	var newGame domain.GameMainInfo
	if err := c.ShouldBindJSON(&newGame); err != nil {
		errorMSg := fmt.Sprintf(
			"[GAME CONTROLLER] Unable to process game: %s", err,
		)
		logger.Logger.Error(errorMSg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[GAME CONTROLLER] Request for create game: %v", newGame,
		),
	)
	gameId, err := g.gameService.CreateGame(newGame)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[GAME CONTROLLER] Error in create game: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[GAME CONTROLLER] Game was created with id: %d",
			gameId,
		),
	)
	response := domain.Response{
		Message: "Game successfully created",
		Data:    map[string]int{"gameId": gameId},
	}
	c.JSON(http.StatusCreated, response)
}
