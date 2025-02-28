package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/alex-arraga/backend_store/pkg/logger"
)

// TODO: secretKey and appEnv should be held from an .env file

const (
	secretKey = "randomString"
	MaxAge    = 86400 * 30 // 30 days
	IsProd    = false
)

func NewAuth() {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if googleClientID == "" || googleClientSecret == "" {
		logger.UseLogger().Fatal().Msg("Missing required environment variables: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET")
	}

	store := sessions.NewCookieStore([]byte(secretKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   MaxAge,
		HttpOnly: true,
		Secure:   IsProd,
	}

	// Assign session storage to Gothic
	gothic.Store = store

	// Config OAuth providers
	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, callbackURL, "email", "profile"),
	)
}
