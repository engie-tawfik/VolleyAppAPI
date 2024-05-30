package ports

import (
	"volleyapp/domain/models"

	"github.com/gin-gonic/gin"
)

type SetRepository interface {
	SaveNewSet(models.SetMainInfo) (int, error)
	FinishSet(int, models.Set) (int, error)
	GetSet(int) (models.Set, error)
	SaveSet(models.Set) (int, error)
}

type SetService interface {
	CreateSet(models.SetMainInfo) (int, error)
	FinishSet(int) (int, error)
	PlaySet(models.Rally) (int, error)
}

type SetController interface {
	CreateSet(*gin.Context)
	FinishSet(*gin.Context)
	InitSetRoutes()
}
