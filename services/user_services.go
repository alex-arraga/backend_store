package services

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/alex-arraga/backend_store/database/gorm_models"
	"github.com/alex-arraga/backend_store/models"
	"github.com/alex-arraga/backend_store/repositories"
	"github.com/alex-arraga/backend_store/utils"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(requestingUserID, targetUserID string, userReq *models.UpdateUser) (*models.UpdateUser, error)
	DeleteUserByID(id string) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	usersDB, err := s.repo.GetAllUsers()
	if err != nil {
		return []models.User{}, fmt.Errorf("couldn't get users: %d", err)
	}

	var allUsers []models.User

	for _, user := range usersDB {
		u := models.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: *user.Password,
		}

		allUsers = append(allUsers, u)
	}

	return allUsers, nil
}

func (s *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	userDB, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("couldn't get user: %d", err)
	}

	userReq := models.User{
		ID:       userDB.ID,
		Name:     userDB.Name,
		Email:    userDB.Email,
		Password: *userDB.Password,
	}

	return &userReq, nil
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

func (s *UserServiceImpl) UpdateUser(requestingUserID, targetUserID string, userReq *models.UpdateUser) (*models.UpdateUser, error) {
	// Get the user that try set changes
	requestingUser, err := s.repo.GetUserByID(requestingUserID)
	if err != nil {
		return nil, err
	}

	// Check if the target user exist
	targetUser, err := s.repo.GetUserByID(targetUserID)
	if err != nil {
		return nil, err
	}

	// Validate if requesting user is admin, just an admin can set different roles
	if userReq.Role != nil && requestingUser.Role != "admin" {
		return nil, fmt.Errorf("to change roles you must be an administrator: %d", err)
	}

	// Change just existing fields in the request
	if userReq.Name != nil {
		targetUser.Name = *userReq.Name
	}
	if userReq.Email != nil {
		targetUser.Email = *userReq.Email
	}
	if userReq.Password != nil {
		hashedPassword, err := utils.HashPassword(*userReq.Password)
		if err != nil {
			return nil, err
		}
		targetUser.Password = &hashedPassword
	}

	// Call repo and apply changes in db
	updatedUser, err := s.repo.UpdateUser(targetUser)
	if err != nil {
		return nil, err
	}

	u := models.UpdateUser{
		Name:     &updatedUser.Name,
		Email:    &updatedUser.Email,
		Role:     &updatedUser.Role,
		Password: updatedUser.Password,
	}

	return &u, nil
}

func (s *UserServiceImpl) DeleteUserByID(id string) error {
	if err := s.repo.DeleteUserByID(id); err != nil {
		return err
	}
	return nil
}
