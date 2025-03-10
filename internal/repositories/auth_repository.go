package repositories

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/pkg/hasher"
)

type AuthRepository interface {
	CreateUser(user *gorm_models.User) (*gorm_models.User, error)
	GetUserByEmail(email string) (*gorm_models.User, error)
	RegisterUserWithEmail(user *gorm_models.User) (*gorm_models.User, error)
	LoginUserWithEmail(email, password string) (*gorm_models.User, error)
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

func (repo *RepoConnection) GetUserByEmail(email string) (*gorm_models.User, error) {
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

func (repo *RepoConnection) RegisterUserWithEmail(user *gorm_models.User) (*gorm_models.User, error) {
	// If user not exist, create them
	newUser, err := repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (repo *RepoConnection) LoginUserWithEmail(email, password string) (*gorm_models.User, error) {
	// Verify if user exist
	userDB, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if userDB.PasswordHash == nil {
		return nil, errors.New("this email is linked to an OAuth account")
	}

	// Verify if password exist
	if err = hasher.CheckPassword(*userDB.PasswordHash, password); err != nil {
		return nil, err
	}

	return userDB, nil
}
