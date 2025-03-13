package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/markbates/goth"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/repositories"
	"github.com/alex-arraga/backend_store/pkg/auth"
	"github.com/alex-arraga/backend_store/pkg/hasher"
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
	LoginWithEmailAndPassword(email, password string) (*models.AuthResponse, error)
	LoginWithOAuth(user goth.User) (*models.UserResponse, error)
}

// * Local Auth services

func (s *authServiceImpl) RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error) {
	// Check if the user exist
	existingUser, _ := s.repo.FindUserByEmail(userReq.Email)
	if existingUser.ID != uuid.Nil {
		return nil, errors.New("user already exists")
	}

	// Create a gorm.User model, the "Provider" field will be created as "local" by default, and "EmailVerified" as "false"
	u := &gorm_models.User{
		ID:           uuid.New(),
		FullName:     userReq.FullName,
		Email:        userReq.Email,
		PasswordHash: &userReq.PasswordHash,
	}

	// Send data to repository
	newUser, err := s.repo.CreateUser(u)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	// Generate client response
	userResp := &models.UserResponse{
		ID:        newUser.ID,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		Role:      newUser.Role,
		Provider:  newUser.Provider,
		AvatarURL: newUser.AvatarURL,
	}

	return userResp, nil
}

func (s *authServiceImpl) LoginWithEmailAndPassword(email, password string) (*models.AuthResponse, error) {
	existingUser, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("error logging in user: %w", err)
	}

	if existingUser.PasswordHash == nil {
		return nil, errors.New("this email is linked to an OAuth account")
	}

	// Verify if password exist
	if err = hasher.CheckPassword(*existingUser.PasswordHash, password); err != nil {
		return nil, err
	}

	// Generate JWT
	token, err := auth.GenerateJWT(existingUser.ID, existingUser.Email)
	if err != nil {
		return &models.AuthResponse{}, err
	}

	// Generate client response
	authResponse := &models.AuthResponse{
		User: models.UserResponse{
			ID:        existingUser.ID,
			FullName:  existingUser.FullName,
			Email:     existingUser.Email,
			Role:      existingUser.Role,
			Provider:  existingUser.Provider,
			AvatarURL: existingUser.AvatarURL,
		},
		Token: token,
	}

	return authResponse, nil
}

// * OAuth services

func (s *authServiceImpl) LoginWithOAuth(user goth.User) (*models.UserResponse, error) {
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
