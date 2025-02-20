package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (port, dbConn string, e error) {
	const envFile string = "../../.env"

	err := godotenv.Load(envFile)
	if err != nil {
		return "", "", fmt.Errorf("rrror loading .env: %d", err)
	}

	port, err = getEnv("PORT")
	if err != nil {
		return "", "", fmt.Errorf("error getting PORT: %d", err)
	}

	dbConn, err = getEnv("DATABASE_URL")
	if err != nil {
		return "", "", fmt.Errorf("error getting DATABASE_URL: %d", err)
	}

	return port, dbConn, nil
}

func getEnv(value string) (string, error) {
	env := os.Getenv(value)
	if env == "" {
		return "", fmt.Errorf("couldn't get env: %s", value)
	}
	return env, nil
}
