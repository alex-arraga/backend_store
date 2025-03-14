package hasher

import "golang.org/x/crypto/bcrypt"

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
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
