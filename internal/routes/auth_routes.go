package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
	"github.com/alex-arraga/backend_store/pkg/jsonutil"
)

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		jsonutil.RespondJSON(w, http.StatusOK, "Auth routes mounted", struct{}{})
	})

	// Inits authentication process with Gothic
	r.Get("/{provider}", func(w http.ResponseWriter, r *http.Request) {
		jsonutil.RespondJSON(w, http.StatusOK, "Provider route", struct{}{})
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
