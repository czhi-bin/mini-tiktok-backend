package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt the password and return the hashed password in string format
func Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// Check if the password is consistent with the hashed password saved in the database
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}