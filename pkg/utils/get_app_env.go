package utils

import "os"

// Represent the possible enviroments where the application can running
type envType string

const (
	envProd  envType = "prod"
	envTest  envType = "test"
	envDev   envType = "dev"
	envLocal envType = "local"
)

// GetAppEnv verifies and returns the enviroment where application is running.
// The possible values of return are: "dev", "test", "prod" or "local"
func GetAppEnv() envType {
	appEnv := os.Getenv("APP_ENV")

	switch envType(appEnv) {
	case envProd, envTest, envDev:
		return envType(appEnv)
	default:
		return envLocal
	}
}
