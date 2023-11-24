package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TeamController struct {
	teamService       ports.TeamService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.TeamController = (*TeamController)(nil)

func NewTeamHandler(
	TeamService ports.TeamService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleware ports.HeadersMiddleware,
) *TeamController {
	domain.RegisterTeamValidators()
	return &TeamController{
		teamService:       TeamService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (t *TeamController) CreateTeam(c *gin.Context) {
	var team domain.NewTeam
	response := domain.Response{
		Message: "",
		Data:    nil,
	}
	// Validate NewTeam data
	if err := c.ShouldBindJSON(&team); err != nil {
		logger.Logger.Error("Unable to process Team")
		var ve validator.ValidationErrors
		var errorMsg string
		// Check for validation errors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				errorMsg = domain.GetTeamErrorMsg(fe)
				if errorMsg != "" {
					break
				}
			}
		}
		response.Message = fmt.Sprintf("Unable to process team. %s", errorMsg)
		logger.Logger.Error("Validation error for team: " + errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	logger.Logger.Info("Request for CreateTeam with email " + team.Email)
	logger.Logger.Debug(fmt.Sprintf("Team data: %+v", team))
	_, err := t.teamService.CreateTeam(team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	response.Message = fmt.Sprintf(
		"Team %s successfully registered",
		team.Name,
	)
	c.JSON(http.StatusCreated, response)
}

func (t *TeamController) GetTeam(c *gin.Context) {
	response := domain.Response{}
	teamId, _ := c.Get("teamId")
	logger.Logger.Info("Request for GetTeam. TeamId: " + teamId.(string))
	team, err := t.teamService.GetTeam(teamId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	response.Message = "Team found"
	response.Data = team
	logger.Logger.Info("Team was found: " + team.Name)
	logger.Logger.Debug(fmt.Sprintf("Team data: %+v", team))
	c.JSON(http.StatusOK, response)
}

func (t *TeamController) UpdateTeamInfo(c *gin.Context) {

}

func (t *TeamController) RegisterTeamRoutes(rg *gin.RouterGroup) {
	teamRoute := rg.Group("/teams", t.headersMiddleware.RequireApiKey)
	teamRoute.GET("", t.authMiddleware.RequireAuth, t.GetTeam)
	teamRoute.POST("", t.CreateTeam)
}
