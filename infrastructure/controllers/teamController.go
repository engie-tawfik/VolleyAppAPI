package controllers

import (
	"fmt"
	"log"
	"net/http"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
	"volleyapp/infrastructure/errors"

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
	models.RegisterTeamValidators()
	return &TeamController{
		gin:               gin,
		teamService:       teamService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (t *TeamController) InitTeamRoutes() {
	teamBasePath := fmt.Sprintf("%s/teams", config.BasePath)
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
	var newTeam models.TeamMainInfo
	if err := c.ShouldBindJSON(&newTeam); err != nil {
		log.Println("Team controller - Unable to process team:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	userId, _ := c.Get("userId")
	newTeam.UserId = int(userId.(float64))
	log.Println("Team controller - CreateTeam request - team:", newTeam)
	teamId, err := t.teamService.CreateTeam(newTeam)
	if err != nil {
		log.Println("Team controller - error in CreateTeam:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	response := models.Response{
		Message: fmt.Sprintf(
			"Team %s successfully registered", newTeam.Name,
		),
		Data: teamId,
	}
	log.Println("Team controller - team created - response:", response)
	c.JSON(http.StatusCreated, response)
}

func (t *TeamController) GetUserTeams(c *gin.Context) {
	userId, _ := c.Get("userId")
	log.Println(
		"Team controller - GetUserTeams request - userId:",
		userId.(float64),
	)
	userTeams, err := t.teamService.GetUserTeams(int(userId.(float64)))
	if err != nil {
		log.Println("Team controller - error in GetUserTeams:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	responseMsg := "Found user teams"
	if len(userTeams) == 0 {
		responseMsg = "No teams found for user"
	}
	response := models.Response{
		Message: responseMsg,
		Data:    map[string][]models.TeamSummary{"userTeams": userTeams},
	}
	log.Println("Team controller - user teams - response:", response)
	c.JSON(http.StatusOK, response)
}

func (t *TeamController) GetTeam(c *gin.Context) {
	response := models.Response{}
	c.JSON(http.StatusOK, response)
}

func (t *TeamController) UpdateTeamInfo(c *gin.Context) {

}
