package services

import (
	"errors"
	"fmt"

	// "github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/repositories"
	"github.com/alex-arraga/backend_store/pkg/hasher"
)

// Implementation and initialization of user services that connect to the user repository
type userServiceImpl struct {
	repo repositories.UserRepository
}

func newUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// Methods of user services
type UserService interface {
	GetAllUsers() ([]models.UserResponse, error)
	FindUserByID(id string) (*models.UserResponse, error)
	UpdateUser(requestingUserID, targetUserID string, dataToUpdate *models.UpdateUser) (*models.UserResponse, error)
	DeleteUserByID(id string) error
}

func (s *userServiceImpl) GetAllUsers() ([]models.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return []models.UserResponse{}, fmt.Errorf("couldn't get users: %w", err)
	}

	var allUsers []models.UserResponse

	for _, user := range users {
		u := models.UserResponse{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Role:      user.Role,
			Provider:  user.Provider,
			AvatarURL: user.AvatarURL,
		}

		allUsers = append(allUsers, u)
	}

	return allUsers, nil
}

func (s *userServiceImpl) FindUserByID(id string) (*models.UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("couldn't get user: %w", err)
	}

	userResponse := &models.UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Provider:  user.Provider,
		AvatarURL: user.AvatarURL,
	}

	return userResponse, nil
}

func (s *userServiceImpl) UpdateUser(requestingUserID, targetUserID string, dataToUpdate *models.UpdateUser) (*models.UserResponse, error) {
	// Get the user that try set changes
	requestingUser, err := s.repo.GetUserByID(requestingUserID)
	if err != nil {
		return nil, fmt.Errorf("requesting user not found: %w", err)
	}

	// Check if the target user exist
	targetUser, err := s.repo.GetUserByID(targetUserID)
	if err != nil {
		return nil, fmt.Errorf("target user not found: %w", err)
	}

	// Validate if requesting user wants to change data of other users
	if requestingUser.Role != "admin" && requestingUserID != targetUserID {
		return nil, errors.New("you must be an administrator to change data of other user")
	}

	// Validate if requesting user wants to change his role for 'admin'
	if dataToUpdate.Role != nil && *dataToUpdate.Role == "admin" && requestingUser.Role != "admin" {
		return nil, errors.New("you must be an administrator to change to 'admin' role")
	}

	// Change just existing fields in the request
	if dataToUpdate.FullName != nil {
		targetUser.FullName = *dataToUpdate.FullName
	}
	if dataToUpdate.Email != nil {
		targetUser.Email = *dataToUpdate.Email
	}
	if dataToUpdate.AvatarURL != nil {
		targetUser.AvatarURL = dataToUpdate.AvatarURL
	}
	if dataToUpdate.Role != nil {
		targetUser.Role = *dataToUpdate.Role
	}
	if dataToUpdate.Password != nil {
		hashedPassword, err := hasher.HashPassword(*dataToUpdate.Password)
		if err != nil {
			return nil, fmt.Errorf("failed hashing password: %w", err)
		}
		targetUser.PasswordHash = &hashedPassword
	}

	// Call repo and apply changes in db
	updatedUser, err := s.repo.UpdateUser(targetUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Return a user response with the new data
	userResp := &models.UserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		Role:      updatedUser.Role,
		Provider:  updatedUser.Provider,
		AvatarURL: updatedUser.AvatarURL,
	}

	return userResp, nil
}

func (s *userServiceImpl) DeleteUserByID(id string) error {
	if err := s.repo.DeleteUserByID(id); err != nil {
		return err
	}
	return nil
}
