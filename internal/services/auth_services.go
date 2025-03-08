package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/markbates/goth"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/repositories"
)

// Implementation and initialization of auth services that connect to the auth repository
type authServiceImpl struct {
	repo repositories.AuthRepository
}

func newAuthService(repo repositories.AuthRepository) AuthServices {
	return &authServiceImpl{repo: repo}
}

// Methods of auth services
type AuthServices interface {
	RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error)
	RegisterWithOAuth(user goth.User) (*models.UserResponse, error)
	LoginWithEmailAndPassword(email, password string) (*models.UserResponse, error)
}

// * Local Auth services

func (s *authServiceImpl) RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error) {
	// Create a gorm.User model, the "Provider" field will be created as "local" by default, and "EmailVerified" as "false"
	u := &gorm_models.User{
		ID:           uuid.New(),
		FullName:     userReq.FullName,
		Email:        userReq.Email,
		PasswordHash: &userReq.PasswordHash,
	}

	// Send data to repository
	userDB, err := s.repo.RegisterUserWithEmail(u)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	userResp := &models.UserResponse{
		ID:        userDB.ID,
		FullName:  userDB.FullName,
		Email:     userDB.Email,
		Role:      userDB.Role,
		Provider:  userDB.Provider,
		AvatarURL: userDB.AvatarURL,
	}

	return userResp, nil
}

func (s *authServiceImpl) LoginWithEmailAndPassword(email, password string) (*models.UserResponse, error) {
	userDB, err := s.repo.LoginUserWithEmail(email, password)
	if err != nil {
		return nil, fmt.Errorf("error logging in user: %w", err)
	}

	userResp := &models.UserResponse{
		ID:        userDB.ID,
		FullName:  userDB.FullName,
		Email:     userDB.Email,
		Role:      userDB.Role,
		Provider:  userDB.Provider,
		AvatarURL: userDB.AvatarURL,
	}

	return userResp, nil
}

// * OAuth services

func (s *authServiceImpl) RegisterWithOAuth(user goth.User) (*models.UserResponse, error) {
	// Converts goth.User (OAuth) to gorm_model.User, in order to be able to send it to the database
	u := gorm_models.User{
		ID:            uuid.New(),
		FullName:      user.Name,
		Email:         user.Email,
		EmailVerified: true,
		Provider:      user.Provider,
		ProviderID:    &user.UserID,
		AvatarURL:     &user.AvatarURL,
	}

	// Send data to database
	userDB, err := s.repo.CreateUser(&u)
	if err != nil {
		return &models.UserResponse{}, err
	}

	// Converts gorm_model.User (database model) to models.UserResponse, in order to be able to send it to the client
	userResponse := models.UserResponse{
		ID:        userDB.ID,
		FullName:  userDB.FullName,
		Email:     userDB.Email,
		Role:      userDB.Role,
		Provider:  userDB.Provider,
		AvatarURL: userDB.AvatarURL,
	}

	return &userResponse, nil
}
