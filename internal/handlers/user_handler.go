package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/services"
	"github.com/alex-arraga/backend_store/pkg/context_keys"
	"github.com/alex-arraga/backend_store/pkg/jsonutil"
)

// path: /user - GET
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	// Get userID from context
	userID, ok := r.Context().Value(context_keys.UserIDKey).(string)
	if !ok {
		jsonutil.RespondError(w, http.StatusUnauthorized, "UserID not found in context")
		return
	}

	// Get user from ID and validate if it's admin
	user, err := us.GetUserByID(userID)
	if err != nil {
		jsonutil.RespondError(w, http.StatusUnauthorized, err.Error())
		return
	}

	if user.Role != "admin" {
		jsonutil.RespondError(w, http.StatusUnauthorized, "Unauthorize: You must be 'admin' to view all users")
		return
	}

	allUses, err := us.GetAllUsers()
	if err != nil {
		jsonutil.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting users: %v", err))
		return
	}

	jsonutil.RespondJSON(w, http.StatusOK, "All users successfully obtained", allUses)
}

// path: /user/{userID} - GET
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "User id is required")
		return
	}

	user, err := us.GetUserByID(userID)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("User with id: %s not exist", userID))
		return
	}

	jsonutil.RespondJSON(w, http.StatusOK, "User successfully found", user)
}

// path /user/{targetUserID} - PUT
func UpdateUserHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	type parameters struct {
		Name      *string `json:"name,omitempty"`
		Email     *string `json:"email,omitempty"`
		Password  *string `json:"password,omitempty"`
		Role      *string `json:"role,omitempty"`
		AvatarURL *string `json:"avatar,omitempty"`
	}

	params, err := jsonutil.ParseRequestBody[parameters](r)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %s", err))
		return
	}

	dataToUpdate := models.UpdateUser{
		FullName:  params.Name,
		Email:     params.Email,
		Password:  params.Password,
		Role:      params.Role,
		AvatarURL: params.AvatarURL,
	}

	// Get UserID from request
	targetUserID := chi.URLParam(r, "targetUserID")

	// Get UserID from context
	requestingUserID, ok := r.Context().Value(context_keys.UserIDKey).(string)
	if !ok || requestingUserID == "" {
		jsonutil.RespondError(w, http.StatusUnauthorized, "Unauthorized: UserID in context not found")
		return
	}

	// Call service
	result, err := us.UpdateUser(requestingUserID, targetUserID, &dataToUpdate)
	if err != nil {
		jsonutil.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating user: %v", err))
	}

	userResponse := models.UserResponse{
		ID:        result.ID,
		FullName:  result.FullName,
		Email:     result.Email,
		Role:      result.Role,
		AvatarURL: result.AvatarURL,
	}

	jsonutil.RespondJSON(w, http.StatusOK, "User updated successfully", userResponse)
}

// path: /user/{userID} - DELETE
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "User id is required")
		return
	}

	if err := us.DeleteUserByID(userID); err != nil {
		jsonutil.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting user: %v", err))
		return
	}

	jsonutil.RespondJSON(w, http.StatusOK, "User deleted successfully", map[string]interface{}{})
}
