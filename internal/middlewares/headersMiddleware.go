package middlewares

import (
	"log"
	"net/http"
	"os"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"

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
		logger.Logger.Error("[HEADERS MIDDLEWARE] Bad headers in request")
		log.Println("Bad headers in request")
		badRequestResponse := domain.Response{
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
	if h.AppArena != os.Getenv("WEBAPP") {
		return false
	}
	if h.HereComesTheBoom != os.Getenv("MOCK_API_KEY") {
		return false
	}
	return true
}
