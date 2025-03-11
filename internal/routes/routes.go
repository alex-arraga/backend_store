package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/alex-arraga/backend_store/internal/middlewares"
	"github.com/alex-arraga/backend_store/internal/services"
)

// Private or Protected routes
func mountProtectedRoutes(r chi.Router, services *services.ServicesContainer) chi.Router {
	r.Group(func(r chi.Router) {
		loadProtectedUserRoutes(r, services.UserSrv)
	})
	return r
}

// Public routes
func mountPublicRoutes(r chi.Router, services *services.ServicesContainer) chi.Router {
	r.Mount("/auth", loadPublicAuthRoutes(services.AuthSrv))
	return r
}

func MountRoutes(services *services.ServicesContainer) chi.Router {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		// Protected routes group
		r.Group(func(r chi.Router) {
			r.Use(middlewares.JWTAuthMiddleware)
			mountProtectedRoutes(r, services)
		})

		// Public routes group
		r.Group(func(r chi.Router) {
			mountPublicRoutes(r, services)
		})
	})

	return r
}
