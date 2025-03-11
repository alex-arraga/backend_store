package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/middlewares"
	"github.com/alex-arraga/backend_store/internal/services"
)

func MountRoutes(services *services.ServicesContainer) chi.Router {
	r := chi.NewRouter()

	v1Router := chi.NewRouter()

	// Use middlewares
	v1Router.Use(middlewares.JWTAuthMiddleware)

	// Mount routes
	v1Router.Mount("/user", loadUserRoutes(services.UserSrv))
	v1Router.Mount("/auth", loadAuthRoutes(services.AuthSrv))

	r.Mount("/v1", v1Router)
	return r
}
