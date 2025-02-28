package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/alex-arraga/backend_store/pkg/logger"
)

type Opts struct {
	GoogleClientID     string
	GoogleClientSecret string
	CallbackURL        string
	SecretKey          string
	MaxAge             int
	HttpOnly           bool
	SecureMode         bool
}

type Config struct {
	opts Opts
}

func loadOptions() Opts {
	// Google opts
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if googleClientID == "" || googleClientSecret == "" || callbackURL == "" {
		logger.UseLogger().Fatal().Msg("Missing required environment variables: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, CALLBACK_URL")
		panic("Missing required environment variables: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, CALLBACK_URL")
	}

	// App enviroment opts
	var isProd bool
	var httpOnly bool

	secretKey := os.Getenv("SECRET_KEY")
	appEnv := os.Getenv("APP_ENV")

	if secretKey == "" || appEnv == "" {
		logger.UseLogger().Fatal().Msg("Missing required enviroment variables: SECRET_KEY or APP_ENV")
		panic("Missing required environment variables: SECRET_KEY or APP_ENV")
	}

	// If application are in "dev" enviroment, isProd will be false, otherwise will be true
	if appEnv == "dev" {
		isProd = false
	} else {
		isProd = true
	}

	// If application are in "dev" enviroment, httpOnly will be true, otherwise will be false
	if isProd {
		httpOnly = false
	} else {
		httpOnly = true
	}

	return Opts{
		GoogleClientID:     googleClientID,
		GoogleClientSecret: googleClientSecret,
		CallbackURL:        callbackURL,
		MaxAge:             86400 * 30,
		SecretKey:          secretKey,
		SecureMode:         isProd,
		HttpOnly:           httpOnly,
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
		google.New(config.opts.GoogleClientID, config.opts.GoogleClientSecret, config.opts.CallbackURL, "email", "profile"),
	)
}
