package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/alex-arraga/backend_store/pkg/logger"
)

type GoogleOpts struct {
	ClientID     string
	ClientSecret string
}

type Opts struct {
	SecretKey   string
	MaxAge      int
	HttpOnly    bool
	SecureMode  bool
	CallbackURL string
	Google      GoogleOpts
}

type Config struct {
	opts Opts
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

func loadOptions() Opts {
	// Load OAuth providers options
	googleOpts := loadGoogleOpts()

	// Application enviroment opts
	var isProd bool
	var httpOnly bool

	secretKey := os.Getenv("SECRET_KEY")
	appEnv := os.Getenv("APP_ENV")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if secretKey == "" || appEnv == "" || callbackURL == "" {
		logger.UseLogger().Fatal().Msg("Missing required enviroment variables: SECRET_KEY, APP_ENV, CALLBACK_URL")
		panic("Missing required environment variables: SECRET_KEY, APP_ENV, CALLBACK_URL")
	}

	// If application is in "dev" enviroment, isProd will be false, otherwise it will be true
	if appEnv == "dev" {
		isProd = false
	} else {
		isProd = true
	}

	// If application is in "dev" enviroment, httpOnly will be true, otherwise it will be false
	if isProd {
		httpOnly = false
	} else {
		httpOnly = true
	}

	return Opts{
		MaxAge:      86400 * 30,
		SecretKey:   secretKey,
		SecureMode:  isProd,
		HttpOnly:    httpOnly,
		CallbackURL: callbackURL,
		Google: GoogleOpts{
			ClientID:     googleOpts.ClientID,
			ClientSecret: googleOpts.ClientSecret,
		},
	}
}

func newConfig() *Config {
	o := loadOptions()

	return &Config{
		opts: o,
	}
}

func NewAuth() {
	config := newConfig()

	store := sessions.NewCookieStore([]byte(config.opts.SecretKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.opts.MaxAge,
		HttpOnly: config.opts.HttpOnly,
		Secure:   config.opts.SecureMode,
	}

	// Assign session storage to Gothic
	gothic.Store = store

	// Config OAuth providers
	goth.UseProviders(
		google.New(config.opts.Google.ClientID, config.opts.Google.ClientSecret, config.opts.CallbackURL, "email", "profile"),
	)
}
