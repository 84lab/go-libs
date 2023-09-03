package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a unique token based on the provided key string.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate from password: %w", err)
	}

	return string(bytes), nil
}

// CheckPassword checks if hashedPassword equals password.
func CheckPassword(hashedPassword, password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
