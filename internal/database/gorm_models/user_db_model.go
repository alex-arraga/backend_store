package gorm_models

import (
	"time"

	"github.com/google/uuid"
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
