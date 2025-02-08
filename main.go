package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/alex-arraga/backend_store/config"
	"github.com/alex-arraga/backend_store/database/connection"
	"github.com/alex-arraga/backend_store/database/migrations"
	"github.com/alex-arraga/backend_store/models"
)

func main() {
	port, dbConn, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %d", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hola mundo"))
		if err != nil {
			log.Println("Error writing response")
		}
	})

	// Database connection
	db, err := connection.ConnectDatabase(dbConn)
	if err != nil {
		log.Fatalf("Database error: %d", err)
	}

	// Execute migrations
	migrations.Migrate(&models.User{})

	// Example: create a user
	user := models.User{
		ID:    uuid.New(),
		Name:  "John",
		Email: "john@example.com",
	}

	if result := db.Create(&user); result.Error != nil {
		log.Fatalf("Error creating user: %d", result.Error)
	}
	log.Printf("User created: %+v\n", user)

	s := http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Print("Server listening...")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed")
	}
}
