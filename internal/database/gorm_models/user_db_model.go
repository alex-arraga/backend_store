package gorm_models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;unique;not null"`
	Password  *string   `gorm:"size:255;default:null"`
	Role      string    `gorm:"size:10;not null;default:user"`
	CreatedAT time.Time `gorm:"autoCreateTime"`
	UpdatedAT time.Time `gorm:"autoUpdateTime"`
}

// Encrypt a password and saved in User.Password
func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*u.Password = string(hash)
	return nil
}

// Compares a password with the hashed password saved
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password))
	return err == nil
}

// Enum of roles
var validRoles = map[string]bool{
	"user":   true,
	"seller": true,
	"admin":  true,
}

// Checks if the user role exist and set them
func (u *User) SetRole(role string) error {
	if !validRoles[role] {
		return fmt.Errorf("invalid role: %s", role)
	}
	return nil
}
