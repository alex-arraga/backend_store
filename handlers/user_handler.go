package handlers

import (
	"fmt"
	"net/http"

	"github.com/alex-arraga/backend_store/models"
	"github.com/alex-arraga/backend_store/services"
	"github.com/alex-arraga/backend_store/utils"
)

// Handler /user - POST
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
