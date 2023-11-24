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
	// Get team based on email
	var user domain.User
	query := "SELECT * FROM users WHERE user_email = $1"
	result, err := a.db.GetDB().Query(
		query,
		email,
	)
	if err != nil {
		return user, fmt.Errorf(
			"[DATABASE] Error in Insert into users: %s", err.Error(),
		)
	}
	if err := result.Scan(
		&user.UserId,
		&user.IsActive,
		&user.Email,
		&user.Password,
		&user.CreationDate,
		&user.LastUpdateDate,
	); err != nil {
		return user, fmt.Errorf(
			"[DATABASE] Error in Insert into users: %s", err.Error(),
		)
	}
	return user, nil
}

func (a *AuthRepository) SaveNewUser(newUser domain.User) (int, error) {
	query := "INSERT INTO users (is_active, user_email, user_password, creation_date, last_update_date) VALUES($1, $2, $3, $4, $5)"
	result, err := a.db.GetDB().Exec(
		query,
		newUser.IsActive,
		newUser.Email,
		newUser.Password,
		newUser.CreationDate,
		newUser.LastUpdateDate,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in Insert into users: %s", err.Error(),
		)
	}
	newUserId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in LastInsertId: %s", err.Error(),
		)
	}
	return int(newUserId), nil
}
