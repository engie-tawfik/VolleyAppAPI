package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type GameRepository interface {
	SaveNewGame(domain.NewGame) (int, error)
}

type GameService interface {
	CreateGame(domain.NewGame) (int, error)
}

type GameController interface {
	CreateGame(g *gin.Context)
	InitGameRoutes()
}
