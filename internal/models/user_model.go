package models

import (
	"github.com/google/uuid"
)

// Models to entry data
type User struct {
	ID            uuid.UUID `json:"id"`
	FullName      string    `json:"fullname"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	PasswordHash  string    `json:"password_hash"`
	Provider      string    `json:"provider"`
}

type UpdateUser struct {
	FullName  *string `json:"fullname,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
	Role      *string `json:"role,omitempty"`
	Provider  *string `json:"provider,omitempty"`
	AvatarURL *string `json:"avatar,omitempty"`
}

// Models to output data
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Provider  string    `json:"provider"`
	AvatarURL *string   `json:"avatar"`
}
