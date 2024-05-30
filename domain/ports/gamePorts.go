package ports

import (
	"volleyapp/domain/models"

	"github.com/gin-gonic/gin"
)

type GameRepository interface {
	SaveNewGame(models.GameMainInfo) (int, error)
	FinishGame(int, models.GameMainInfo) (int, error)
	GetTeamsNames(int) (models.GameTeamsNames, error)
	GetGame(int) (models.Game, error)
	SaveGame(models.Game) (int, error)
}

type GameService interface {
	CreateGame(models.GameMainInfo) (int, error)
	FinishGame(int) (int, error)
	UpdateGameStats(models.Set) (int, error)
}

type GameController interface {
	CreateGame(*gin.Context)
	FinishGame(*gin.Context)
	InitGameRoutes()
}
