package repositories

import (
	"fmt"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"
)

type AuthRepository struct {
	db ports.Database
}

var _ ports.AuthRepository = (*AuthRepository)(nil)

func NewAuthRepository(db ports.Database) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE user_email = $1"
	result := a.db.GetDB().QueryRow(
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
			"[DATABASE] Error in get user by email: %s", err,
		)
	}
	return user, nil
}

func (a *AuthRepository) SaveNewUser(newUser domain.User) (int, error) {
	query := "INSERT INTO users (is_active, user_email, user_password, creation_date, last_update_date) VALUES($1, $2, $3, $4, $5) RETURNING user_id"
	result := a.db.GetDB().QueryRow(
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
			"[DATABASE] Error in save new user: %s", err,
		)
	}
	return int(newUserId), nil
}
