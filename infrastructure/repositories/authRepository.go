package repositories

import (
	"fmt"
	"volleyapp/config"
	"volleyapp/domain/models"
	"volleyapp/domain/ports"
)

type AuthRepository struct{}

var _ ports.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (a *AuthRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE user_email = $1"
	result := config.DB.QueryRow(
		query,
		email,
	)
	if err := result.Scan(
		&user.UserId,
		&user.IsActive,
		&user.Email,
		&user.Password,
		&user.CreationDate,
		&user.LastUpdateDate,
	); err != nil {
		return user, fmt.Errorf(
			"database - error in GetUserByEmail: %v", err,
		)
	}
	return user, nil
}

func (a *AuthRepository) SaveNewUser(newUser models.User) (int, error) {
	query := "INSERT INTO users (is_active, user_email, user_password, creation_date, last_update_date) VALUES($1, $2, $3, $4, $5) RETURNING user_id"
	result := config.DB.QueryRow(
		query,
		newUser.IsActive,
		newUser.Email,
		newUser.Password,
		newUser.CreationDate,
		newUser.LastUpdateDate,
	)
	var newUserId int
	err := result.Scan(&newUserId)
	if err != nil {
		return 0, fmt.Errorf(
			"database - error in SaveNewUser: %v", err,
		)
	}
	return int(newUserId), nil
}
