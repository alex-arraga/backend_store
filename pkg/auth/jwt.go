package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey []byte

type Claims struct {
	UserID uuid.UUID
	Email  string
	jwt.RegisteredClaims
}

func loadJWTKey() error {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		return errors.New("couldn't get enviroment variable: JWT_KEY")
	}

	// Converts string to []bytes
	keyBytes := ([]byte)(key)

	jwtKey = keyBytes
	return nil
}

func GenerateJWT(userID uuid.UUID, email string) (string, error) {
	if err := loadJWTKey(); err != nil {
		return "", errors.New(err.Error())
	}

	expirationTime := time.Now().Add(168 * time.Hour) // 1 week
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Validate JWT
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
