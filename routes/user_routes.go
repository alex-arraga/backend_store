package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/models"
	"github.com/alex-arraga/backend_store/utils"
)

func userRoutes() chi.Router {
	r := chi.NewRouter()

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondJSON(w, http.StatusOK, "Received a 'GET' request in /user route")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		params, err := utils.ParseRequestBody[parameters](r)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %s", err))
		}

		if params.Name == "" {
			utils.RespondError(w, http.StatusBadRequest, "Name is required")
		}
		if params.Email == "" {
			utils.RespondError(w, http.StatusBadRequest, "Email is required")
		}
		if params.Password == "" {
			utils.RespondError(w, http.StatusBadRequest, "Password is required")
		}

		userReq := models.User{
			Name:     params.Name,
			Email:    params.Email,
			Password: params.Password,
		}

		fmt.Print(userReq)

		utils.RespondJSON(w, http.StatusOK, "Received a 'POST' request in /user route")
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondError(w, http.StatusOK, "Received a ERROR in 'PUT' request in /user route")
	})

	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondError(w, http.StatusOK, "Received a ERROR in 'DELETE' request in /user route")
	})

	return r
}
