package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %d \n", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing JSON: %d \n", err)
	}
}

func RespondError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error: %s \n", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondJSON(w, code, errResponse{Error: msg})
}

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&params); err != nil {
		return params, err
	}
	return params, nil
}
