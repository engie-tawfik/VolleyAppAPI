package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var unauthorizedResponse = domain.Response{
	ErrorCode: http.StatusUnauthorized,
	Message:   "Unauthorized.",
	Data:      nil,
}

type AuthMiddleware struct {
}

var _ ports.AuthMiddleware = (*AuthMiddleware)(nil)

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) RequireAuth(c *gin.Context) {
	// Get cookie from request
	tokenString, err := c.Cookie("Access")
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[AUTH MIDDLEWARE] Error getting Access cookie: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}

	// Decode token
	var secretBytes = []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm for token is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				error := fmt.Errorf(
					"[AUTH MIDDLEWARE] Unexpected signing method: %v",
					token.Header["alg"],
				)
				return nil, error
			}
			return secretBytes, nil
		})
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[AUTH MIDDLEWARE] Error parsing token: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		logger.Logger.Error("[AUTH MIDDLEWARE] Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}
	c.Next()
}

func (a *AuthMiddleware) RequireRefresh(c *gin.Context) {
	// Get cookie from request
	tokenString, err := c.Cookie("Refresh")
	if err != nil {
		errorMsg := fmt.Sprintf("[AUTH MIDDLEWARE] Error getting Refresh cookie: %s", err)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}

	// Decode token
	var secretBytes = []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm for token is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretBytes, nil
	})
	if err != nil {
		log.Println("Error parsing token:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		log.Println("Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}
	c.Next()
}
