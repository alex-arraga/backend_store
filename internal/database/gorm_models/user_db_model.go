package gorm_models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FullName      string    `gorm:"size:255;default:null"`
	Email         string    `gorm:"uniqueIndex;not null"`
	EmailVerified bool      `gorm:"default:false"`
	PasswordHash  *string   `gorm:"size:255;default:null"`
	Provider      string    `gorm:"default:local"`
	ProviderID    *string
	AvatarURL     *string
	Role          string    `gorm:"size:10;not null;default:user"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
