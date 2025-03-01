package auth

import (
	"os"

	"github.com/alex-arraga/backend_store/pkg/logger"
)

type GoogleOpts struct {
	ClientID     string
	ClientSecret string
}

func loadGoogleOpts() *GoogleOpts {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if googleClientID == "" || googleClientSecret == "" {
		logger.UseLogger().Fatal().Msg("Missing required environment variables: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET")
		panic("Missing required environment variables: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET")
	}

	return &GoogleOpts{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
	}
}