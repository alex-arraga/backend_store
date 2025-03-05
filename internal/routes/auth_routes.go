package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
	"github.com/alex-arraga/backend_store/internal/services"
)

func authRoutes(as services.AuthServices) chi.Router {
	r := chi.NewRouter()

	// Register using email and password
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUserHandler(w, r, as)
	})

	// Starts OAuth login
	r.Get("/{provider}/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.BeginAuthLoginHandler(w, r)
	})

	// Receives Google response and get the authenticated user
	r.Get("/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAuthCallbackHandler(w, r, as)
	})

	return r
}
