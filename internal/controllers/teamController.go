package controllers

import (
	"fmt"
	"net/http"
	"os"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/errors"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	gin               *gin.Engine
	teamService       ports.TeamService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.TeamController = (*TeamController)(nil)

func NewTeamController(
	gin *gin.Engine,
	teamService ports.TeamService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleware ports.HeadersMiddleware,
) *TeamController {
	domain.RegisterTeamValidators()
	return &TeamController{
		gin:               gin,
		teamService:       teamService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (t *TeamController) InitTeamRoutes() {
	teamBasePath := fmt.Sprintf("%s/teams", os.Getenv("BASE_PATH"))
	teamRoute := t.gin.Group(
		teamBasePath,
		t.headersMiddleware.RequireApiKey,
		t.authMiddleware.RequireAuth,
	)
	teamRoute.POST("/create", t.CreateTeam)
	teamRoute.GET("/user", t.GetUserTeams)
	teamRoute.GET("/team", t.GetTeam)
}

func (t *TeamController) CreateTeam(c *gin.Context) {
	var newTeam domain.TeamMainInfo
	if err := c.ShouldBindJSON(&newTeam); err != nil {
		errorMSg := fmt.Sprintf(
			"[TEAM CONTROLLER] Unable to process team: %s", err,
		)
		logger.Logger.Error(errorMSg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	userId, _ := c.Get("userId")
	newTeam.UserId = int(userId.(float64))
	logger.Logger.Info(
		fmt.Sprintf(
			"[TEAM CONTROLLER] Request for create team: %s", newTeam.Name,
		),
	)
	logger.Logger.Debug(
		fmt.Sprintf(
			"[TEAM CONTROLLER] Team data: %+v", newTeam))
	teamId, err := t.teamService.CreateTeam(newTeam)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[TEAM CONTROLLER] Error in create team: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[TEAM CONTROLLER] Team was created with id: %d",
			teamId,
		),
	)
	response := domain.Response{
		Message: fmt.Sprintf(
			"Team %s successfully registered", newTeam.Name,
		),
		Data: teamId,
	}
	c.JSON(http.StatusCreated, response)
}

func (t *TeamController) GetUserTeams(c *gin.Context) {
	userId, _ := c.Get("userId")
	logger.Logger.Info(
		fmt.Sprintf(
			"[TEAM CONTROLLER] Request for get user teams: %v", userId.(float64),
		),
	)
	userTeams, err := t.teamService.GetUserTeams(int(userId.(float64)))
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[TEAM CONTROLLER] Error in get user teams: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	responseMsg := "Found user teams"
	if len(userTeams) == 0 {
		responseMsg = "No teams found for user"
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[TEAM CONTROLLER] User teams: %v",
			userTeams,
		),
	)
	response := domain.Response{
		Message: responseMsg,
		Data:    map[string][]domain.TeamSummary{"userTeams": userTeams},
	}
	c.JSON(http.StatusOK, response)
}

func (t *TeamController) GetTeam(c *gin.Context) {
	response := domain.Response{}
	c.JSON(http.StatusOK, response)
}

func (t *TeamController) UpdateTeamInfo(c *gin.Context) {

}
