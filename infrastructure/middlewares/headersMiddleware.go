package middlewares

import (
	"log"
	"net/http"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"

	"github.com/gin-gonic/gin"
)

type HeadersMiddleware struct {
	HereComesTheBoom string `header:"Here-Comes-The-Boom" binding:"required"`
	AppArena         string `header:"App-Arena" binding:"required"`
}

var _ ports.HeadersMiddleware = (*HeadersMiddleware)(nil)

func NewHeadersMiddleware() *HeadersMiddleware {
	return &HeadersMiddleware{}
}

func (h *HeadersMiddleware) RequireApiKey(c *gin.Context) {
	var headers HeadersMiddleware
	err := c.ShouldBindHeader(&headers)
	if err != nil || !passValidations(headers) {
		log.Println("Headers middleware - bad headers in request")
		badRequestResponse := models.Response{
			ErrorCode: http.StatusBadRequest,
			Message:   "Bad request",
			Data:      nil,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, badRequestResponse)
		return
	}
	c.Next()
}

func passValidations(h HeadersMiddleware) bool {
	if h.HereComesTheBoom == "" || h.AppArena == "" {
		return false
	}
	if h.AppArena != config.WebApp {
		return false
	}
	if h.HereComesTheBoom != config.ApiKey {
		return false
	}
	return true
}
