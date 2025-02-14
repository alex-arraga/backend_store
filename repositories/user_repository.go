package repositories

import (
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/database/gorm_models"
)

type UserRepository interface {
	GetAllUsers() ([]gorm_models.User, error)
	GetUserByID(id string) (*gorm_models.User, error)
	CreateUser(user *gorm_models.User) error
}

type RepoConnection struct {
	db *gorm.DB
}

func newUserRepo(db *gorm.DB) UserRepository {
	return &RepoConnection{db: db}
}

func (repo *RepoConnection) CreateUser(user *gorm_models.User) error {
	return repo.db.Create(user).Error
}

func (repo *RepoConnection) GetUserByID(id string) (*gorm_models.User, error) {
	var user gorm_models.User

	result := repo.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *RepoConnection) GetAllUsers() ([]gorm_models.User, error) {
	var users []gorm_models.User

	result := repo.db.Find(&users)
	if result.Error != nil {
		return []gorm_models.User{}, result.Error
	}

	return users, nil
}
