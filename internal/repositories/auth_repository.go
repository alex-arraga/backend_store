package repositories

import (
	"errors"

	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/pkg/hasher"
)

type AuthRepository interface {
	CreateUser(user *gorm_models.User) (*gorm_models.User, error)
	GetUserByEmail(email string) (*gorm_models.User, error)
	RegisterUserWithEmail(user *gorm_models.User) (*gorm_models.User, error)
	LoginUserWithEmail(email, passwordHashed string) (*gorm_models.User, error)
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

	userDB := repo.db.First(&user, "email = ?", email)
	if userDB.Error != nil {
		return nil, errors.New("user not found in database")
	}

	return &user, nil
}

func (repo *RepoConnection) RegisterUserWithEmail(user *gorm_models.User) (*gorm_models.User, error) {
	// Verify if user exist
	userDB, _ := repo.GetUserByEmail(user.Email)

	// If user exist, execute login
	if userDB != nil {
		if _, err := repo.LoginUserWithEmail(user.Email, *user.PasswordHash); err != nil {
			return nil, err
		}
	}

	// If user not exist, create them
	u, err := repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// TODO: Generate and return a JWT

	return u, nil
}

func (repo *RepoConnection) LoginUserWithEmail(email, passwordHashed string) (*gorm_models.User, error) {
	// Verify if user exist
	userDB, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if userDB.PasswordHash == nil {
		return nil, errors.New("this email is linked to an OAuth account")
	}

	// Verify if password exist
	if err = hasher.CheckPassword(*userDB.PasswordHash, passwordHashed); err != nil {
		return nil, err
	}

	// TODO: Generate and return a JWT

	return userDB, nil
}
