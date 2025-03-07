package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/alex-arraga/backend_store/internal/models"
	"github.com/alex-arraga/backend_store/internal/services"
	"github.com/alex-arraga/backend_store/pkg/hasher"
	"github.com/alex-arraga/backend_store/pkg/jsonutil"
	"github.com/alex-arraga/backend_store/pkg/logger"
	"github.com/alex-arraga/backend_store/pkg/utils"
)

// * Local register handler

// Register in the local application using an email and password.
// Path to call: /v1/auth/register
func RegisterUserWithEmailHandler(w http.ResponseWriter, r *http.Request, as services.AuthServices) {
	type parameters struct {
		FullName string `json:"fullname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params, err := jsonutil.ParseRequestBody[parameters](r)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %s", err))
		return
	}

	if params.FullName == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "FullName is required")
		return
	}
	if params.Email == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "Email is required")
		return
	}
	if params.Password == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "Password is required")
		return
	}

	// Hashing password
	hashedPassword, err := hasher.HashPassword(params.Password)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, "Error hashing password")
		return
	}

	u := models.User{
		FullName:     params.FullName,
		Email:        params.Email,
		PasswordHash: hashedPassword,
	}

	user, err := as.RegisterWithEmailAndPassword(&u)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	jsonutil.RespondJSON(w, http.StatusOK, "User successfully registered", user)
}

// Login in the local application using an email and password.
// Path to call: /v1/auth/login
func LoginUserWithEmailHandler(w http.ResponseWriter, r *http.Request, as services.AuthServices) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params, err := jsonutil.ParseRequestBody[parameters](r)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid input: %s", err))
		return
	}

	if params.Email == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "Email is required")
		return
	}
	if params.Password == "" {
		jsonutil.RespondError(w, http.StatusBadRequest, "Password is required")
		return
	}

	// Hashing password
	hashedPassword, err := hasher.HashPassword(params.Password)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, "Error hashing password")
		return
	}

	user, err := as.LoginWithEmailAndPassword(params.Email, hashedPassword)
	if err != nil {
		jsonutil.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonutil.RespondJSON(w, http.StatusOK, "User successfully registered", user)
}

// * OAuth handlers

// Handler to starting OAuth login.
// Path to call: /v1/auth/{provider}/login?provider={provider}.
// Example path: /v1/auth/{google}/login?provider=google
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

// Handler to managaments Google respond after authentication.
// Path to call: /v1/auth/{provider}/login?provider={provider}.
// Example path: /v1/auth/google/login?provider=google
func GetAuthCallbackHandler(w http.ResponseWriter, r *http.Request, as services.AuthServices) {
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

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		logger.UseLogger().Error().Str("module", "handlers").Str("nameFunc", "GetAuthCallback").Str("error", err.Error()).Msg("OAuth authentication failed")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	logger.UseLogger().Debug().Msgf("User authenticated \n %v", gothUser)

	// Send user data to User Service
	user, err := as.RegisterWithOAuth(gothUser)
	if err != nil {
		jsonutil.RespondError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user with OAuth: %v", err))
		return
	}

	// Respond with JSON to check data if "APP_ENV" is dev
	if utils.GetAppEnv() != "prod" {
		jsonutil.RespondJSON(w, http.StatusOK, "User successfully registered with OAuth", user)
		return
	}

	// Redirect when auth successfully
	clientURL := os.Getenv("CLIENT_REDIRECT_URL")
	if clientURL == "" {
		logger.UseLogger().Error().Str("module", "handlers").Str("var", "clientURL").Msg("Missing clientURL env")
		jsonutil.RespondError(w, http.StatusInternalServerError, "Missing clientURL env")
		return
	}

	http.Redirect(w, r, clientURL, http.StatusFound)
}
