package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
)

func authRoutes() chi.Router {
	r := chi.NewRouter()

	// Inits authentication process with Gothic
	r.Get("/{provider}/login}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAuthCallback(w, r)
	})

	// Receives Google response and get the authenticated user
	r.Get("/{provider}/callback}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAuthCallback(w, r)
	})

	return r
}
