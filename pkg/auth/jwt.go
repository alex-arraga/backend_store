package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/alex-arraga/backend_store/pkg/logger"
)

var jwtKey []byte

type Claims struct {
	UserID string
	Email  string
	jwt.RegisteredClaims
}

func LoadJWTKey() {
	key := os.Getenv("JWT_KEY")
	if key == "" {
		logger.UseLogger().Fatal().Msg("couldn't get enviroment variable: JWT_KEY")
	}

	// Converts string to []bytes
	keyBytes := ([]byte)(key)

	jwtKey = keyBytes
}

func GenerateJWT(userID uuid.UUID, email string) (string, error) {
	expirationTime := time.Now().Add(168 * time.Hour) // 1 week
	claims := &Claims{
		UserID: userID.String(),
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

	// Validate
	if _, err := uuid.Parse(claims.UserID); err != nil {
		return nil, errors.New("invalid UUID format in UserID")
	}

	return claims, nil
}
