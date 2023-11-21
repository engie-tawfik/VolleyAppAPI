package ports

import "github.com/gin-gonic/gin"

type HeadersMiddleware interface {
	RequireApiKey(c *gin.Context)
}
