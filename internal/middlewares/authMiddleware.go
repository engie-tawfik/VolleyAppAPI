package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"

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
	unauthorizedResponse := domain.Response{
		Message: "Unauthorized.",
		Data:    nil,
	}

	// Get cookie from request
	tokenString, err := c.Cookie("Access")
	if err != nil {
		log.Println("Error getting Access cookie:", err)
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
		c.Set("teamId", claims["sub"])
	} else {
		log.Println("Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}
	c.Next()
}

func (a *AuthMiddleware) RequireRefresh(c *gin.Context) {
	unauthorizedResponse := domain.Response{
		Message: "Unauthorized.",
		Data:    nil,
	}

	// Get cookie from request
	tokenString, err := c.Cookie("Refresh")
	if err != nil {
		log.Println("Error getting Access cookie:", err)
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
		c.Set("teamId", claims["sub"])
	} else {
		log.Println("Invalid token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
		return
	}
	c.Next()
}
