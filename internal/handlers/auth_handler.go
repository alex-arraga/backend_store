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

	// Recover session before to complete authentication
	sess, err := gothic.Store.Get(r, "auth-session")
	if err != nil {
		logger.UseLogger().Error().Str("module", "handlers").Str("error", err.Error()).Msg("Failed to get session")
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	// Ensure that the provider in saved session is the same as the URL
	savedProvider, ok := sess.Values["provider"].(string)
	if !ok || savedProvider != provider {
		logger.UseLogger().Error().Str("module", "handlers").Str("nameFunc", "GetAuthCallback").Msg("Provider mismatch in session")
		http.Error(w, "Provider mismatch in session", http.StatusUnauthorized)
		return
	}

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		logger.UseLogger().Error().Str("module", "handlers").Str("nameFunc", "GetAuthCallback").Str("error", err.Error()).Msg("OAuth authentication failed")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	logger.UseLogger().Info().Msgf("User authenticated \n %v", user)

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

	// Get session and save provider
	sess, err := gothic.Store.Get(r, "auth-session")
	if err != nil {
		logger.UseLogger().Error().Str("module", "handlers").Str("error", err.Error()).Msg("Failed to get session")
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	sess.Values["provider"] = provider
	err = sess.Save(r, w)
	if err != nil {
		logger.UseLogger().Error().Str("module", "handlers").Str("error", err.Error()).Msg("Failed to save session")
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	logger.UseLogger().Debug().Msgf("Starting authentication with %s ", provider)

	// BeginAuthHandler expects received a query parameter with the provider. Ex: localhost:8000/auth?provider=google
	// then automatically redirects to Google for finish user authentication
	gothic.BeginAuthHandler(w, r)
}
