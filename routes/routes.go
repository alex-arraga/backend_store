package routes

import (
	"github.com/go-chi/chi/v5"
)

func MountRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Use(middleware.Auth)

	v1Router := chi.NewRouter()
	v1Router.Mount("/user", userRoutes())

	r.Mount("/v1", v1Router)
	return r
}
