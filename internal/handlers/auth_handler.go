package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/alex-arraga/backend_store/pkg/jsonutil"
	"github.com/alex-arraga/backend_store/pkg/logger"
)

// Handler to managaments Google respond after authentication
func GetAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
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
	clientURL := os.Getenv("CLIENT_REDIRECT_URL")
	if clientURL == "" {
		logger.UseLogger().Error().Str("module", "handlers").Str("var", "clientURL").Msg("Missing clientURL env")
		jsonutil.RespondError(w, http.StatusInternalServerError, "Missing clientURL env")
		return
	}

	http.Redirect(w, r, clientURL, http.StatusFound)
}

// Handler to starting OAuth login
func BeginAuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		logger.UseLogger().Error().Str("module", "handlers").Str("nameFunc", "AuthLoginHandler").Msg("Missing provider in OAuth process")
		jsonutil.RespondError(w, http.StatusBadRequest, "Missing provider")
		return
	}

	logger.UseLogger().Info().Msgf("Starting authentication with %s ", provider)

	// BeginAuthHandler expects received a query parameter with the provider. Ex: localhost:8000/auth?provider=google
	// then automatically redirects to Google for finish user authentication
	gothic.BeginAuthHandler(w, r)
}
