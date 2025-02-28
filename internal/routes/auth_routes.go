package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
)

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Route /auth - OK")
	})

	// Inits authentication process with Gothic
	r.Get("/{provider}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Route /auth/{provider} - OK")
	})

	r.Get("/{provider}/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.BeginAuthLoginHandler(w, r)
	})

	// Receives Google response and get the authenticated user
	r.Get("/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAuthCallbackHandler(w, r)
	})

	return r
}
