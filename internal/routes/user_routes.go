package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
	"github.com/alex-arraga/backend_store/internal/services"
)

func userRoutes(us services.UserService) chi.Router {
	r := chi.NewRouter()

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllUsersHandler(w, r, us)
	})

	r.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserByIDHandler(w, r, us)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// handlers.CreateUserHandler(w, r, us)
	})

	r.Put("/{targetUserID}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(w, r, us)
	})

	r.Delete("/{userID}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUserHandler(w, r, us)
	})

	return r
}
