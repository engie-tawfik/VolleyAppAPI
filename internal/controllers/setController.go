package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/errors"
	"volleyapp/logger"

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
	return &SetController{
		gin:               gin,
		setService:        setService,
		authMiddleware:    authMiddleware,
		headersMiddleware: headersMiddleware,
	}
}

func (s *SetController) InitSetRoutes() {
	setBasePath := fmt.Sprintf("%s/sets", os.Getenv("BASE_PATH"))
	setRoute := s.gin.Group(
		setBasePath,
		s.headersMiddleware.RequireApiKey,
		s.authMiddleware.RequireAuth,
	)
	setRoute.POST("/create", s.CreateSet)
	setRoute.PUT("/finish/:setId", s.FinishSet)
}

func (s *SetController) CreateSet(c *gin.Context) {
	var newSet domain.SetMainInfo
	if err := c.ShouldBindJSON(&newSet); err != nil {
		errorMSg := fmt.Sprintf(
			"[SET CONTROLLER] Unable to process set: %s", err,
		)
		logger.Logger.Error(errorMSg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[SET CONTROLLER] Request for create set: %v", newSet,
		),
	)
	setId, err := s.setService.CreateSet(newSet)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[SET CONTROLLER] Error in create set: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[SET CONTROLLER] Set was created with id: %d",
			setId,
		),
	)
	response := domain.Response{
		Message: "Set successfully started",
		Data:    map[string]int{"setId": setId},
	}
	c.JSON(http.StatusCreated, response)
}

func (s *SetController) FinishSet(c *gin.Context) {
	setId, err := strconv.ParseInt(c.Param("setId"), 10, 64)
	if err != nil {
		errorMSg := fmt.Sprintf(
			"[SET CONTROLLER] Unable to process set id: %s", err,
		)
		logger.Logger.Error(errorMSg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[SET CONTROLLER] Request for finish set: %v", setId,
		),
	)
	rowsAffected, err := s.setService.FinishSet(int(setId))
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[SET CONTROLLER] Error in finish set: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestResponse)
		return
	}
	logger.Logger.Info(
		fmt.Sprintf(
			"[SET CONTROLLER] Set was finished with id: %d - %d rows affected",
			setId,
			rowsAffected,
		),
	)
	response := domain.Response{
		Message: "Set successfully finished",
		Data:    nil,
	}
	c.JSON(http.StatusCreated, response)
}
