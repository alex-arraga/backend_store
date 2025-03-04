package services

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/repositories"
	"github.com/alex-arraga/backend_store/pkg/hasher"
)

type UserService interface {
	GetAllUsers() ([]models.UserResponse, error)
	GetUserByID(id string) (*models.UserResponse, error)
	RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error)
	// TODO: LoginWithOAuth(userReq *models.User) (*models.UserResponse, error)
	UpdateUser(requestingUserID, targetUserID string, userReq *models.UpdateUser) (*models.UserResponse, error)
	DeleteUserByID(id string) error
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.UserResponse, error) {
	usersDB, err := s.repo.GetAllUsers()
	if err != nil {
		return []models.UserResponse{}, fmt.Errorf("couldn't get users: %w", err)
	}

	var allUsers []models.UserResponse

	for _, user := range usersDB {
		u := models.UserResponse{
			ID: user.ID,
			// Name:  user.FullName,
			Email: user.Email,
			Role:  user.Role,
		}

		allUsers = append(allUsers, u)
	}

	return allUsers, nil
}

func (s *UserServiceImpl) GetUserByID(id string) (*models.UserResponse, error) {
	userDB, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("couldn't get user: %w", err)
	}

	userReq := models.UserResponse{
		ID: userDB.ID,
		// Name:  userDB.FullName,
		Email: userDB.Email,
		Role:  userDB.Role,
	}

	return &userReq, nil
}

func (s *UserServiceImpl) RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error) {
	// Hashing password
	hashedPassword, err := hasher.HashPassword(userReq.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	u := &gorm_models.User{
		ID:       uuid.New(),
		FullName: userReq.FullName,
		Email:    userReq.Email,
		// EmailVerified: false,
		PasswordHash: &hashedPassword,
		// Provider: "local",
	}

	// Send data to repository
	userDB, err := s.repo.CreateUser(u)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	userResp := &models.UserResponse{
		ID:       userDB.ID,
		FullName: userDB.FullName,
		Email:    userDB.Email,
		Role:     userDB.Role,
		Provider: userDB.Provider,
	}

	return userResp, nil
}

func (s *UserServiceImpl) UpdateUser(requestingUserID, targetUserID string, userReq *models.UpdateUser) (*models.UserResponse, error) {
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

	// Validate if requesting user is admin, just an admin can set different roles
	if userReq.Role != nil && requestingUser.Role != "admin" {
		return nil, fmt.Errorf("you must be an administrator to change roles")
	}

	// Change just existing fields in the request
	// if userReq.Name != nil {
	// 	targetUser.FullName = *userReq.Name
	// }
	// if userReq.Email != nil {
	// 	targetUser.Email = *userReq.Email
	// }
	// if userReq.Password != nil {
	// 	hashedPassword, err := hasher.HashPassword(*userReq.Password)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed hashing password: %w", err)
	// 	}
	// 	targetUser.PasswordHash = &hashedPassword
	// }

	// Call repo and apply changes in db
	updatedUser, err := s.repo.UpdateUser(targetUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	u := models.UserResponse{
		ID: updatedUser.ID,
		// Name:  updatedUser.FullName,
		Email: updatedUser.Email,
		Role:  updatedUser.Role,
	}

	return &u, nil
}

func (s *UserServiceImpl) DeleteUserByID(id string) error {
	if err := s.repo.DeleteUserByID(id); err != nil {
		return err
	}
	return nil
}
