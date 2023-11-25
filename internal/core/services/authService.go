package services

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
	"volleyapp/logger"
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

func (a *AuthService) Login(
	email, password string,
) (domain.AuthResponse, error) {
	var response domain.AuthResponse
	user, err := a.authRepository.GetUserByEmail(email)
	if err != nil {
		errorMsg := fmt.Sprintf("[AUTH SERVICE] Error in gogin: %s", err)
		return response, fmt.Errorf(errorMsg)
	}

	// Verify password
	if passOk := utils.Verify(password, user.Password); !passOk {
		errorMsg := "[AUTH SERVICE] Error in login: can't hash password"
		return response, fmt.Errorf(errorMsg)
	}

	response, err = a.CreateTokens(user.UserId)
	if err != nil {
		errorMsg := fmt.Sprintf("[AUTH SERVICE] Error in login: %s", err)
		return response, fmt.Errorf(errorMsg)
	}

	return response, nil
}

func (a *AuthService) CreateTokens(userId int) (domain.AuthResponse, error) {
	var response domain.AuthResponse
	var tokenLife = os.Getenv("JWT_TOKEN_EXPIRE_MINUTES")
	var secretBytes = []byte(os.Getenv("SECRET"))
	// Create tokens
	tokenLifeInt, err := strconv.Atoi(tokenLife)
	if err != nil {
		errorMsg := fmt.Sprintf(
			"[AUTH SERVICE] Error in create tokens: %s", err,
		)
		return response, fmt.Errorf(errorMsg)
	}
	accessTokenLifeDuration := time.Duration(tokenLifeInt) * time.Minute
	refreshTokenLifeDuration := time.Duration(tokenLifeInt*2) * time.Minute
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(accessTokenLifeDuration).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(refreshTokenLifeDuration).Unix(),
	})
	accessTokenString, _ := accessToken.SignedString(secretBytes)
	refreshTokenString, _ := refreshToken.SignedString(secretBytes)
	log.Println("Access token:", accessTokenString)
	log.Println("Refresh token:", refreshTokenString)
	response = domain.AuthResponse{
		AccessToken:  accessTokenString,
		Refreshtoken: refreshTokenString,
	}
	return response, nil
}

func (a *AuthService) CreateUser(newUser domain.User) (int, error) {
	hashedPass := utils.Hash(newUser.Password)
	newUser.Password = hashedPass
	newUser.IsActive = true
	loc, _ := time.LoadLocation("America/Bogota")
	newUser.CreationDate = time.Now().In(loc)
	newUser.LastUpdateDate = time.Now().In(loc)
	userId, err := a.authRepository.SaveNewUser(newUser)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s", err))
		return 0, fmt.Errorf(
			"[AUTH SERVICE] Error in create user: %s", err,
		)
	}
	return userId, nil
}
