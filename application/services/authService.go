package services

import (
	"fmt"
	"log"
	"time"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
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
) (models.AuthResponse, error) {
	var response models.AuthResponse
	user, err := a.authRepository.GetUserByEmail(email)
	if err != nil {
		return response, fmt.Errorf("auth service - error in Login: %v", err)
	}

	// Verify password
	if passOk := utils.Verify(password, user.Password); !passOk {
		return response, fmt.Errorf(
			"auth service - error in Login: can't hash password",
		)
	}

	response, err = a.CreateTokens(user.UserId)
	if err != nil {
		return response, fmt.Errorf("auth service - error in Login: %v", err)
	}

	return response, nil
}

func (a *AuthService) CreateTokens(userId int) (models.AuthResponse, error) {
	var response models.AuthResponse
	// Create tokens
	accessTokenLifeDuration :=
		time.Duration(config.JwtExpireMins) * time.Minute
	refreshTokenLifeDuration :=
		time.Duration(config.JwtExpireMins*2) * time.Minute
	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(accessTokenLifeDuration).Unix(),
		},
	)
	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(refreshTokenLifeDuration).Unix(),
		},
	)
	accessTokenString, _ := accessToken.SignedString(config.Secret)
	refreshTokenString, _ := refreshToken.SignedString(config.Secret)
	log.Println("Access token:", accessTokenString)
	log.Println("Refresh token:", refreshTokenString)
	response = models.AuthResponse{
		AccessToken:  accessTokenString,
		Refreshtoken: refreshTokenString,
	}
	return response, nil
}

func (a *AuthService) CreateUser(newUser models.User) (int, error) {
	hashedPass := utils.Hash(newUser.Password)
	newUser.Password = hashedPass
	newUser.IsActive = true
	loc, _ := time.LoadLocation("America/Bogota")
	newUser.CreationDate = time.Now().In(loc)
	newUser.LastUpdateDate = time.Now().In(loc)
	userId, err := a.authRepository.SaveNewUser(newUser)
	if err != nil {
		return 0, fmt.Errorf("auth service - error in CreateUser: %v", err)
	}
	return userId, nil
}
