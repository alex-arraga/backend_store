package gorm_models

import (
	"time"

	"github.com/google/uuid"
)

type AuthAccount struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;index"`
	Provider      string    `gorm:"not null"`
	ProviderID    *string   `gorm:"size:255;default:null"`
	Email         string    `gorm:"size:255;not null;index"`
	EmailVerified bool      `gorm:"default:false"`
	PasswordHash  *string   `gorm:"size:255;default:null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
