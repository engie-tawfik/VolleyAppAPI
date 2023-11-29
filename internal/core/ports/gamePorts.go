package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type GameRepository interface {
	SaveNewGame(domain.GameMainInfo) (int, error)
}

type GameService interface {
	CreateGame(domain.GameMainInfo) (int, error)
}

type GameController interface {
	CreateGame(*gin.Context)
	InitGameRoutes()
}
