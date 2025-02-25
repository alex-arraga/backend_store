package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/handlers"
)

func authRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{provider}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAuthCallback(w, r)
	})

	return r
}
