package main

import (
	"net/http"
	"time"

	"github.com/alex-arraga/backend_store/internal/repositories"
	"github.com/alex-arraga/backend_store/internal/routes"
	"github.com/alex-arraga/backend_store/internal/services"
	"github.com/alex-arraga/backend_store/pkg/logger"
)

// StartServer configura y arranca el servidor HTTP.
func StartServer(config *AppConfig) {
	// Cargar repositorios y servicios
	repos := repositories.LoadRepositories(config.DB)
	services := services.LoadServices(repos)

	// Montar rutas
	r := routes.MountRoutes(services)

	// Configurar servidor HTTP
	s := &http.Server{
		Handler:      r,
		Addr:         ":" + config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.UseLogger().Info().Str("port", config.Port).Msgf("Server listening...")
	err := s.ListenAndServe()
	if err != nil {
		logger.UseLogger().Fatal().Err(err).Msg("Server failed")
	}
}
