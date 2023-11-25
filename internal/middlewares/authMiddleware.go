package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/errors"
	"volleyapp/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

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
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Decode token
	var secretBytes = []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm for token is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := fmt.Errorf(
					"[AUTH MIDDLEWARE] Unexpected signing method: %v",
					token.Header["alg"],
				)
				return nil, err
			}
			return secretBytes, nil
		})
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[AUTH MIDDLEWARE] Error parsing token: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		logger.Logger.Error("[AUTH MIDDLEWARE] Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Decode token
	var secretBytes = []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm for token is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf(
				"[AUTH MIDDLEWARE] Unexpected signing method: %v",
				token.Header["alg"],
			)
			return nil, err
		}
		return secretBytes, nil
	})
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[AUTH MIDDLEWARE] Error parsing token: %s", err,
		)
		logger.Logger.Error(errorMsg)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		logger.Logger.Error("[AUTH MIDDLEWARE] Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}
	c.Next()
}
