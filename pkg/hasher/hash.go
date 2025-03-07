package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHash []byte

// Encrypt a password and saved in User.Password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(PasswordHash(hash)), nil
}

// Compares a password with the hashed password saved
func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	return err
}
