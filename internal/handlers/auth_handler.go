package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alex-arraga/backend_store/pkg/jsonutil"
	"github.com/alex-arraga/backend_store/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

type contextKey string

const providerKey contextKey = "provider"

func GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Print(user)

	// Redirect when auth successfully
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

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