package repositories

import (
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAllUsers() ([]gorm_models.User, error)
	FindUserByID(id string) (*gorm_models.User, error)
	FindUserByEmail(email string) (*gorm_models.User, error)
	CreateUser(user *gorm_models.User) (*gorm_models.User, error)
	UpdateUser(user *gorm_models.User) (*gorm_models.User, error)
	DeleteUserByID(id string) error
}

func newUserRepo(db *gorm.DB) UserRepository {
	return &RepoConnection{db: db}
}

func (repo *RepoConnection) GetAllUsers() ([]gorm_models.User, error) {
	var users []gorm_models.User

	if result := repo.db.Find(&users); result.Error != nil {
		return []gorm_models.User{}, result.Error
	}

	return users, nil
}

func (repo *RepoConnection) FindUserByID(id string) (*gorm_models.User, error) {
	// TODO: Validate if id is an UUID
	var user gorm_models.User

	if result := repo.db.First(&user, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
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

func (repo *RepoConnection) CreateUser(user *gorm_models.User) (*gorm_models.User, error) {
	if result := repo.db.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *RepoConnection) UpdateUser(user *gorm_models.User) (*gorm_models.User, error) {
	result := repo.db.Model(&gorm_models.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	updatedUser, err := repo.FindUserByID(user.ID.String())
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (repo *RepoConnection) DeleteUserByID(id string) error {
	var user gorm_models.User

	if result := repo.db.Delete(&user, "id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}
