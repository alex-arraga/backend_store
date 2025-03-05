package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/markbates/goth"

	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/repositories"
)

type AuthServices interface {
	_RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error)
	_RegisterWithOAuth(user goth.User) (*models.UserResponse, error)
}

type authServiceImpl struct {
	repo repositories.AuthRepository
}

func newAuthService(repo repositories.AuthRepository) AuthServices {
	return &authServiceImpl{repo: repo}
}

func (s *authServiceImpl) _RegisterWithEmailAndPassword(userReq *models.User) (*models.UserResponse, error) {
	u := &gorm_models.User{
		ID:       uuid.New(),
		FullName: userReq.FullName,
		Email:    userReq.Email,
		// EmailVerified: false,
		PasswordHash: &userReq.PasswordHash,
		// Provider: "local",
	}

	// Send data to repository
	userDB, err := s.repo.RegisterUser(u)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
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

func (s *authServiceImpl) _RegisterWithOAuth(user goth.User) (*models.UserResponse, error) {
	fullName := fmt.Sprint(user.Name + " " + user.LastName)

	// Converts goth.User (OAuth) to gorm_model.User, in order to be able to send it to the database
	u := gorm_models.User{
		ID:            uuid.New(),
		FullName:      fullName,
		Email:         user.Email,
		EmailVerified: true,
		Provider:      user.Provider,
		ProviderID:    &user.UserID,
		AvatarURL:     &user.AvatarURL,
	}

	// Send data to database
	userDB, err := s.repo.RegisterUser(&u)
	if err != nil {
		return &models.UserResponse{}, fmt.Errorf("")
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
