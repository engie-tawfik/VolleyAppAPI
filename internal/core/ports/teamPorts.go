package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type TeamRepository interface {
	CreateTeam(team domain.NewTeam) (bool, error)
	GetTeam(teamId string) (domain.Team, error)
	CheckTeamExistence(email string) (bool, error)
}

type TeamService interface {
	CreateTeam(domain.NewTeam) (bool, error)
	GetTeam(string) (domain.Team, error)
	UpdateTeamInfo(domain.BaseTeam) (bool, error)
}

type TeamController interface {
	CreateTeam(c *gin.Context)
	GetTeam(c *gin.Context)
	UpdateTeamInfo(c *gin.Context)
	RegisterTeamRoutes(rg *gin.RouterGroup)
}
