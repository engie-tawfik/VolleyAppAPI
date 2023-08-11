package services

import (
	"log"
	"os"
	"strconv"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/utils"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	authRepository ports.AuthRepository
}

var _ ports.AuthService = (*AuthService)(nil)

func NewAuthService(repository ports.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: repository,
	}
}

func (a *AuthService) Login(email, password string) domain.AuthResponse {
	var response domain.AuthResponse
	team := a.authRepository.Login(email)
	if team.Password == "" {
		log.Println("Wrong login credentials")
		return response
	}
	// Verify password
	if passOk := utils.Verify(password, team.Password); !passOk {
		log.Println("Wrong login credentials")
		return response
	}
	return a.CreateTokens(team.TeamId.Hex())
}

func (a *AuthService) CreateTokens(teamId string) domain.AuthResponse {
	var tokenLife = os.Getenv("JWT_TOKEN_EXPIRE_MINUTES")
	var secretBytes = []byte(os.Getenv("SECRET"))
	// Create tokens
	tokenLifeInt, err := strconv.Atoi(tokenLife)
	if err != nil {
		log.Println("Error parsing tokenLife:", err)
	}
	accessTokenLifeDuration := time.Duration(tokenLifeInt) * time.Minute
	refreshTokenLifeDuration := time.Duration(tokenLifeInt*2) * time.Minute
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": teamId,
		"exp": time.Now().Add(accessTokenLifeDuration).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": teamId,
		"exp": time.Now().Add(refreshTokenLifeDuration).Unix(),
	})
	accessTokenString, _ := accessToken.SignedString(secretBytes)
	refreshTokenString, _ := refreshToken.SignedString(secretBytes)
	log.Println("Access token:", accessTokenString)
	log.Println("Refresh token:", refreshTokenString)
	response := domain.AuthResponse{
		AccessToken:  accessTokenString,
		Refreshtoken: refreshTokenString,
	}
	return response
}
