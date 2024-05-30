package models

import (
	"time"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	UserId         int       `json:"userId"`
	IsActive       bool      `json:"isActive"`
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" binding:"required,passwordcheck"`
	CreationDate   time.Time `json:"creationDate"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
}

var PasswordCheck validator.Func = func(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpace := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasNumber = true
		} else if unicode.IsSpace(char) {
			hasSpace = true
		}
	}
	if ok {
		if len(password) >= 12 && hasUpper && hasLower && hasNumber && !hasSpace {
			return true
		}
	}
	return false
}

// Registers custom validators in models for JSON binding
func RegisterUserValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("passwordcheck", PasswordCheck)
	}
}
