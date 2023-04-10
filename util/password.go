package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}
	return string(hashpassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}