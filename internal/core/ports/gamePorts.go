package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type GameRepository interface {
	SaveNewGame(domain.GameMainInfo) (int, error)
	FinishGame(int, domain.GameMainInfo) (int, error)
	GetTeamsNames(int) (domain.GameTeamsNames, error)
	GetGame(int) (domain.Game, error)
	SaveGame(domain.Game) (int, error)
}

type GameService interface {
	CreateGame(domain.GameMainInfo) (int, error)
	FinishGame(int) (int, error)
	UpdateGameStats(domain.Set) (int, error)
}

type GameController interface {
	CreateGame(*gin.Context)
	FinishGame(*gin.Context)
	InitGameRoutes()
}
