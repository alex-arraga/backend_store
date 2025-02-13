package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/handlers"
	"github.com/alex-arraga/backend_store/services"
	"github.com/alex-arraga/backend_store/utils"
)

func userRoutes(us services.UserService) chi.Router {
	r := chi.NewRouter()

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Validate if user role is admin
		utils.RespondJSON(w, http.StatusOK, "Received a 'GET' request in /user route")
	})

	r.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		if userID == "" {
			utils.RespondError(w, http.StatusBadRequest, "User id is required")
		}

		utils.RespondJSON(w, http.StatusOK, "Received a 'GET' request in /user/{i} route: "+userID)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUserHandler(w, r, us)
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondError(w, http.StatusOK, "Received a ERROR in 'PUT' request in /user route")
	})

	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondError(w, http.StatusOK, "Received a ERROR in 'DELETE' request in /user route")
	})

	return r
}
