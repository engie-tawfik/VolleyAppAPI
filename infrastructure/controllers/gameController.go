package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
	"volleyapp/infrastructure/errors"

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
	gameBasePath := fmt.Sprintf("%s/games", config.BasePath)
	gameRoute := g.gin.Group(
		gameBasePath,
		g.headersMiddleware.RequireApiKey,
		g.authMiddleware.RequireAuth,
	)
	gameRoute.POST("/create", g.CreateGame)
	gameRoute.PUT("/finish/:gameId", g.FinishGame)
}

func (g *GameController) CreateGame(c *gin.Context) {
	var newGame models.GameMainInfo
	if err := c.ShouldBindJSON(&newGame); err != nil {
		log.Println("Game controller - unable to process game:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Game controller - CreateGame request - game:", newGame)
	gameId, err := g.gameService.CreateGame(newGame)
	if err != nil {
		log.Println("Game controller - error in CreateGame:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Game controller - game was created with id:")
	response := models.Response{
		Message: "Game successfully created",
		Data:    map[string]int{"gameId": gameId},
	}
	c.JSON(http.StatusCreated, response)
}

func (g *GameController) FinishGame(c *gin.Context) {
	gameId, err := strconv.ParseInt(c.Param("gameId"), 10, 64)
	if err != nil {
		log.Println("Game controller - unable to process game id:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Game controller - FinishGame request - gameId:", gameId)
	rowsAffected, err := g.gameService.FinishGame(int(gameId))
	if err != nil {
		log.Println("Game controller - error in FinishGame:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println(
		"Game controller - game was finished with id:",
		gameId,
		"- rows affected:",
		rowsAffected,
	)
	response := models.Response{
		Message: "Game successfully finished",
		Data:    nil,
	}
	c.JSON(http.StatusCreated, response)
}
