package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/backend_store/config"
	"github.com/alex-arraga/backend_store/database/connection"
	"github.com/alex-arraga/backend_store/database/migrations"
	"github.com/alex-arraga/backend_store/repositories"
	"github.com/alex-arraga/backend_store/routes"
	"github.com/alex-arraga/backend_store/services"
)

func main() {
	port, dbConn, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %d", err)
	}

	// Database connection
	db, err := connection.ConnectDatabase(dbConn)
	if err != nil {
		log.Fatalf("Database error: %d", err)
	}

	// Execute migrations
	migrations.ExecMigrations()
	repos := repositories.LoadRepositories(db)
	services.LoadServices(repos)

	// Load routes
	r := routes.MountRoutes()

	s := http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Print("Server listening...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %d", err)
	}
}
