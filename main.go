package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hola mundo"))
		if err != nil {
			log.Println("Error writing response")
		}
	})

	s := http.Server{
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed")
	}
}
