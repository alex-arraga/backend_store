package services

import (
	"fmt"

	"github.com/alex-arraga/backend_store/database/gorm_models"
	"github.com/alex-arraga/backend_store/models"
	"github.com/alex-arraga/backend_store/repositories"
	"github.com/alex-arraga/backend_store/utils"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	// GetUserByID(id string) (*models.User, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) CreateUser(userReq *models.User) error {
	// Hashing password
	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %d", err)
	}

	user := &gorm_models.User{
		ID:       uuid.New(),
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: &hashedPassword,
	}

	// Send data to repository
	return s.repo.CreateUser(user)
}
