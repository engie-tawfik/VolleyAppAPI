package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"volleyapp/config"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TeamHandler struct {
	teamService       ports.TeamService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.TeamHandler = (*TeamHandler)(nil)

func NewTeamHandler(
	TeamService ports.TeamService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleware ports.HeadersMiddleware,
) *TeamHandler {
	domain.RegisterTeamValidators()
	return &TeamHandler{
		teamService:       TeamService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (t *TeamHandler) CreateTeam(c *gin.Context) {
	var team domain.NewTeam
	response := domain.Response{
		Message: "",
		Data:    nil,
	}
	// Validate NewTeam data
	if err := c.ShouldBindJSON(&team); err != nil {
		config.Logger.Error("Unable to process Team")
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
		config.Logger.Error("Validation error for team: " + errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	config.Logger.Info("Request for CreateTeam with email " + team.Email)
	config.Logger.Debug(fmt.Sprintf("Team data: %+v", team))
	result := t.teamService.CreateTeam(team)
	if result != "" {
		response.Message = result
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = fmt.Sprintf("Team %s successfully registered.", team.Name)
	c.JSON(http.StatusCreated, response)
}

func (t *TeamHandler) GetTeam(c *gin.Context) {
	response := domain.Response{}
	teamId, _ := c.Get("teamId")
	config.Logger.Info("Request for GetTeam. TeamId: " + teamId.(string))
	team := t.teamService.GetTeam(teamId.(string))
	if team.Name == "" {
		response.Message = "Team not found."
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "Team found."
	response.Data = team
	config.Logger.Info("Team was found: " + team.Name)
	config.Logger.Debug(fmt.Sprintf("Team data: %+v", team))
	c.JSON(http.StatusOK, response)
}

func (t *TeamHandler) RegisterTeamRoutes(rg *gin.RouterGroup) {
	teamRoute := rg.Group("/teams", t.headersMiddleware.RequireApiKey)
	teamRoute.GET("", t.authMiddleware.RequireAuth, t.GetTeam)
	teamRoute.POST("", t.CreateTeam)
}
