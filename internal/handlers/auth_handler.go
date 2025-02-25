package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/alex-arraga/backend_store/pkg/jsonutil"
	"github.com/alex-arraga/backend_store/pkg/logger"
)

// Handler to managaments Google respond after authentication
func GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		logger.UseLogger().Error().Str("module", "handlers").Str("nameFunc", "GetAuthCallback").Msg("Missing provider in OAuth process")
		jsonutil.RespondError(w, http.StatusBadRequest, "Missing provider")
		return
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	logger.UseLogger().Info().Msgf("User authenticated %v", user)

	// Redirect when auth successfully
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

// Handler to starting OAuth login
func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		logger.UseLogger().Error().Msg("Missing provider in OAuth process")
		jsonutil.RespondError(w, http.StatusBadRequest, "Missing provider")
		return
	}

	logger.UseLogger().Info().Msgf("Starting authentication with %s ", provider)

	// Automatically redirects to Google for user authentication
	gothic.BeginAuthHandler(w, r)
}
