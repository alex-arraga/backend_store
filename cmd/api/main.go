package main

import (
	"github.com/alex-arraga/backend_store/pkg/logger"
)

func main() {
	config, err := LoadAppConfig()
	if err != nil {
		logger.UseLogger().Fatal().Err(err).Msg("Error loading app config")
	}

	StartServer(config)
}
