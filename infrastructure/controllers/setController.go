package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
	"volleyapp/infrastructure/errors"

	"github.com/gin-gonic/gin"
)

type SetController struct {
	gin               *gin.Engine
	setService        ports.SetService
	authMiddleware    ports.AuthMiddleware
	headersMiddleware ports.HeadersMiddleware
}

var _ ports.SetController = (*SetController)(nil)

func NewSetController(
	gin *gin.Engine,
	setService ports.SetService,
	authMiddleware ports.AuthMiddleware,
	headersMiddleware ports.HeadersMiddleware,
) *SetController {
	models.RegisterSetValidators()
	return &SetController{
		gin:               gin,
		setService:        setService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (s *SetController) InitSetRoutes() {
	setBasePath := fmt.Sprintf("%s/sets", config.BasePath)
	setRoute := s.gin.Group(
		setBasePath,
		s.headersMiddleware.RequireApiKey,
		s.authMiddleware.RequireAuth,
	)
	setRoute.POST("/create", s.CreateSet)
	setRoute.PUT("/finish/:setId", s.FinishSet)
	setRoute.POST("/play", s.PlaySet)
}

func (s *SetController) CreateSet(c *gin.Context) {
	var newSet models.SetMainInfo
	if err := c.ShouldBindJSON(&newSet); err != nil {
		log.Println("Set controller - unable to process set:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Set controller - CreateSet request - set:", newSet)
	setId, err := s.setService.CreateSet(newSet)
	if err != nil {
		log.Println("Set controller - error in CreateSet:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Set controller - set was created with id:", setId)
	response := models.Response{
		Message: "Set successfully started",
		Data:    map[string]int{"setId": setId},
	}
	c.JSON(http.StatusCreated, response)
}

func (s *SetController) FinishSet(c *gin.Context) {
	setId, err := strconv.ParseInt(c.Param("setId"), 10, 64)
	if err != nil {
		log.Println("Set controller - unable to process set id:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Set controller - FinishSet request - setId:", setId)
	rowsAffected, err := s.setService.FinishSet(int(setId))
	if err != nil {
		log.Println("Set controller - error in FinishSet:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println(
		"Set controller - set was finished with id:",
		setId,
		"- rows affected:",
		rowsAffected,
	)
	response := models.Response{
		Message: "Set successfully finished",
		Data:    nil,
	}
	c.JSON(http.StatusCreated, response)
}

func (s *SetController) PlaySet(c *gin.Context) {
	var rally models.Rally
	if err := c.ShouldBindJSON(&rally); err != nil {
		log.Println("Set controller - unable to process rally:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println("Set controller - PlaySet request - rally:", rally)
	rowsAffected, err := s.setService.PlaySet(rally)
	if err != nil {
		log.Println("Set controller - error in PlaySet:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	log.Println(
		"Set controller - rally was saved for set with id:",
		rally.SetId,
		"- rows affected:",
		rowsAffected,
	)
	response := models.Response{
		Message: "Rally successfully saved",
		Data:    nil,
	}
	c.JSON(http.StatusCreated, response)
}
