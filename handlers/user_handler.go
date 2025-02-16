package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/models"
	"github.com/alex-arraga/backend_store/services"
	"github.com/alex-arraga/backend_store/utils"
)

// path: /user - GET
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	// TODO: Validate if user role is admin
	// if user.Role != "admin" {
	// 	return utils.RespondError(w, http.StatusUnauthorized, "Invalid user role")
	// }

	allUses, err := us.GetAllUsers()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting users: %d", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, allUses)
}

// path: /user/{userID} - GET
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		utils.RespondError(w, http.StatusBadRequest, "User id is required")
		return
	}

	user, err := us.GetUserByID(userID)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, fmt.Sprintf("User with id: %s not exist", userID))
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// path: /user - POST
func CreateUserHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params, err := utils.ParseRequestBody[parameters](r)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %s", err))
		return
	}

	if params.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name is required")
		return
	}
	if params.Email == "" {
		utils.RespondError(w, http.StatusBadRequest, "Email is required")
		return
	}
	if params.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, "Password is required")
		return
	}

	userReq := models.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}

	err = us.CreateUser(&userReq)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Error creating user: %d", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, "User created successfully")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request, us services.UserService) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		utils.RespondError(w, http.StatusBadRequest, "User id is required")
		return
	}

	if err := us.DeleteUser(userID); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error deleting user: %d", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, "User deleted successfully")
}
