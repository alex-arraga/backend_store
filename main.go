package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/alex-arraga/backend_store/database/connection"
	"github.com/alex-arraga/backend_store/database/migrations"
	"github.com/alex-arraga/backend_store/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env: %d", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hola mundo"))
		if err != nil {
			log.Println("Error writing response")
		}
	})

	port := os.Getenv("PORT")

	// Database connection
	db, err := connection.ConnectDatabase()
	if err != nil {
		log.Fatalf("Database error: %d", err)
	}

	// Execute migrations
	migrations.Migrate(&models.User{})

	// Example: create a user
	user := models.User{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
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
