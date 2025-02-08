package main

import (
	"log"
	"net/http"
	"os"
	"time"

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
	connection.ConnectDatabase()
	migrations.Migrate(&models.User{
		Name:  "John",
		Email: "john@gmail.com",
	})

	s := http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed")
	}
}
