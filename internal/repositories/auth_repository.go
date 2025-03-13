package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
)

type AuthRepository interface {
	CreateUser(user *gorm_models.User) (*gorm_models.User, error)
	FindUserByEmail(email string) (*gorm_models.User, error)
	FindUserByID(id string) (*gorm_models.User, error)
}

func newAuthRepo(db *gorm.DB) AuthRepository {
	return &RepoConnection{db: db}
}

func (repo *RepoConnection) CreateUser(user *gorm_models.User) (*gorm_models.User, error) {
	if result := repo.db.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *RepoConnection) FindUserByEmail(email string) (*gorm_models.User, error) {
	var user gorm_models.User

	err := repo.db.First(&user, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &gorm_models.User{ID: uuid.Nil}, nil
		}
		return &gorm_models.User{}, err
	}

	return &user, nil
}

func (repo *RepoConnection) FindUserByID(id string) (*gorm_models.User, error) {
	// TODO: Validate if id is an UUID
	var user gorm_models.User

	if result := repo.db.First(&user, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
