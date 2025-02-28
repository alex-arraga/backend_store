package jsonutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alex-arraga/backend_store/pkg/logger"
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
		logger.UseLogger().Error().Err(err).Msg("Failed to marshal JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Config headers and write the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		logger.UseLogger().Warn().Err(err).Msg("Error writing JSON response")
	}
}

func RespondError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		logger.UseLogger().Error().Msgf("Responding with 5XX error: %s \n", msg)
	} else if code >= 400 {
		logger.UseLogger().Warn().Msgf("Responding with 4XX error: %s", msg)
	}

	// Estructura de error con "error" en lugar de "msg"
	response := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}

	// Convertir a JSON
	data, err := json.Marshal(response)
	if err != nil {
		logger.UseLogger().Error().Err(err).Msg("Failed to marshal JSON error response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Configurar headers y escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(data)
	if err != nil {
		logger.UseLogger().Warn().Err(err).Msg("Error writing error JSON response")
	}
}

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	defer r.Body.Close()

	if err := decoder.Decode(&params); err != nil {
		return params, fmt.Errorf("error decoding parameters: %w", err)
	}
	return params, nil
}
