package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func userRoutes() chi.Router {
	r := chi.NewRouter()

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Received a 'GET' request in /user route"))
		if err != nil {
			log.Print("Error in GET user route")
		}
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Received a 'POST' request in /user route"))
		if err != nil {
			log.Print("Error in POST user route")
		}
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Received a 'PUT' request in /user route"))
		if err != nil {
			log.Print("Error in PUT user route")
		}
	})

	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Received a 'DELETE' request in /user route"))
		if err != nil {
			log.Print("Error in DELETE user route")
		}
	})

	return r
}
