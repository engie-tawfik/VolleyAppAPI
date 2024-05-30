package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"volleyapp/config"
	"volleyapp/domain/ports"
	"volleyapp/infrastructure/errors"

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
		log.Println("Auth middleware - error getting access cookie:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Decode token
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm for token is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := fmt.Errorf(
					"auth middleware - unexpected signing method: %v",
					token.Header["alg"],
				)
				return nil, err
			}
			return config.Secret, nil
		})
	if err != nil {
		log.Println("Auth middleware - error parsing token:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		log.Println("Auth middleware - invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}
	c.Next()
}

func (a *AuthMiddleware) RequireRefresh(c *gin.Context) {
	// Get cookie from request
	tokenString, err := c.Cookie("Refresh")
	if err != nil {
		log.Println("Auth middleware - error getting refresh cookie:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm for token is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf(
				"auth middleware - unexpected signing method: %v",
				token.Header["alg"],
			)
			return nil, err
		}
		return config.Secret, nil
	})
	if err != nil {
		log.Println("Auth middleware - error parsing token:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}

	// Get token data
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Attach teamId to request
		c.Set("userId", claims["sub"])
	} else {
		log.Println("Auth middleware - invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.UnauthorizedResponse)
		return
	}
	c.Next()
}
