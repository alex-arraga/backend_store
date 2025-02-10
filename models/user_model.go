package models

import (
	"github.com/google/uuid"
)

// TODO: The API would cans update the user rol
type User struct {
	ID uuid.UUID `json:"id"`
	// Role     string    `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
