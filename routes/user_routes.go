package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func userRoutes() chi.Router {
	r := chi.NewRouter()	

	// /user
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Received a 'GET' request in /user route"))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Received a 'POST' request in /user route"))
	})

	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Received a 'PUT' request in /user route"))
	})

	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Received a 'DELETE' request in /user route"))
	})

	return r
}
