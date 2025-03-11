package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
	"github.com/alex-arraga/backend_store/internal/services"
)

// Public auth routes
func loadPublicAuthRoutes(r chi.Router, as services.AuthServices) chi.Router {
	r.Route("/auth", func(r chi.Router) {
		// Path of group: /v1/auth
		r.Group(func(r chi.Router) {
			// Register using email and password
			r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
				handlers.RegisterUserWithEmailHandler(w, r, as)
			})

			// Login using email and password
			r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
				handlers.LoginUserWithEmailHandler(w, r, as)
			})

			// Starts OAuth login
			r.Get("/{provider}/login", func(w http.ResponseWriter, r *http.Request) {
				handlers.BeginAuthLoginHandler(w, r)
			})

			// Receives Google response and get the authenticated user
			r.Get("/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
				handlers.GetAuthCallbackHandler(w, r, as)
			})
		})
	})

	return r
}
