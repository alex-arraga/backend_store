package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/utils"
)

func userRoutes() chi.Router {
	r := chi.NewRouter()

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondJSON(w, http.StatusOK, "Received a 'GET' request in /user route")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
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
