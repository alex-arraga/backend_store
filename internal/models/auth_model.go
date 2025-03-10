package models

import "github.com/golang-jwt/jwt/v5"

type AuthResponse struct {
	User  UserResponse
	Token jwt.Token
}
