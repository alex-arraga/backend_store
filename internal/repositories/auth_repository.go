package repositories

import (
	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user *gorm_models.User) (*gorm_models.User, error)
	// LoginUser(user *gorm_models.User) (*gorm_models.User, error)
}

func newAuthRepo(db *gorm.DB) AuthRepository {
	return &RepoConnection{db: db}
}

func (repo *RepoConnection) RegisterUser(user *gorm_models.User) (*gorm_models.User, error) {
	if result := repo.db.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}