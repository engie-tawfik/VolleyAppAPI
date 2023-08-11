package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type TeamRepository interface {
	CreateTeam(team domain.NewTeam) (string, bool)
	GetTeam(teamId string) domain.Team
	CheckTeamExistence(email string) (bool, error)
}

type TeamService interface {
	CreateTeam(domain.NewTeam) string
	GetTeam(string) domain.Team
}

type TeamHandler interface {
	CreateTeam(c *gin.Context)
	GetTeam(c *gin.Context)
	RegisterTeamRoutes(rg *gin.RouterGroup)
}
