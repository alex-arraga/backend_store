package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
	"github.com/alex-arraga/backend_store/internal/services"
)

func loadProtectedUserRoutes(r chi.Router, us services.UserService) chi.Router {
	r.Route("/user", func(r chi.Router) {
		// Path of group: /v1/user
		r.Group(func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				handlers.GetAllUsersHandler(w, r, us)
			})

			r.Get("/{userID}", func(w http.ResponseWriter, r *http.Request) {
				handlers.GetUserByIDHandler(w, r, us)
			})

			r.Put("/{targetUserID}", func(w http.ResponseWriter, r *http.Request) {
				handlers.UpdateUserHandler(w, r, us)
			})

			r.Delete("/{userID}", func(w http.ResponseWriter, r *http.Request) {
				handlers.DeleteUserHandler(w, r, us)
			})
		})
	})

	return r
}
