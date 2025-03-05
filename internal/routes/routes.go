package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/services"
)

func MountRoutes(services *services.ServicesContainer) chi.Router {
	r := chi.NewRouter()

	v1Router := chi.NewRouter()
	v1Router.Mount("/user", userRoutes(services.UserSrv))

	// Auth route
	v1Router.Mount("/auth", authRoutes(services.AuthSrv))

	r.Mount("/v1", v1Router)
	return r
}
