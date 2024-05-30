package ports

import (
	"volleyapp/domain/models"

	"github.com/gin-gonic/gin"
)

type TeamRepository interface {
	SaveNewTeam(models.TeamMainInfo) (int, error)
	GetUserTeams(int) ([]models.TeamSummary, error)
	GetTeam(string) (models.Team, error)
	CheckTeamExistence(string) (bool, error)
}

type TeamService interface {
	CreateTeam(models.TeamMainInfo) (int, error)
	GetUserTeams(int) ([]models.TeamSummary, error)
	GetTeam(string) (models.Team, error)
	UpdateTeamInfo(models.TeamMainInfo) (bool, error)
}

type TeamController interface {
	CreateTeam(*gin.Context)
	GetUserTeams(*gin.Context)
	GetTeam(*gin.Context)
	UpdateTeamInfo(*gin.Context)
	InitTeamRoutes()
}
