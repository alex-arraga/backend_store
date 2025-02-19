package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, code int, msg string, payload interface{}) {
	if payload == nil {
		payload = struct{}{}
	}

	// Response struct
	response := struct {
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Msg:  msg,
		Data: payload,
	}

	// Convert to JSON
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Config headers and write the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing JSON response: %d \n", err)
	}
}

func RespondError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("Responding with 5xx error: %s\n", msg)
	}

	// Estructura de error con "error" en lugar de "msg"
	response := map[string]interface{}{
		"error": msg,
	}

	// Convertir a JSON
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal JSON error response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Configurar headers y escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error writing error JSON response: %v\n", err)
	}
}

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()

	if err := decoder.Decode(&params); err != nil {
		return params, fmt.Errorf("error decoding parameters: %v", err)
	}
	return params, nil
}
