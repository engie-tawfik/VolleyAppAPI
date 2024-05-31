package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Hash : Encrypt the user password into a slice of bytes and
// and return a string of the converted bytes
func Hash(password string) string {
	var hashedPass string
	genPassword, err :=
		bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password")
		return hashedPass
	}
	hashedPass = string(genPassword)
	return hashedPass
}

// Verify : this helps to verify the input password while logging in
// and the previously hashed password
func Verify(password, hashedPassword string) bool {
	if password == "" || hashedPassword == "" {
		return false
	}
	err :=
		bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Invalid password comparison error")
			return false
		}
		return false
	}
	return true
}
