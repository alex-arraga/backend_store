package main

import (
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/config"
	"github.com/alex-arraga/backend_store/internal/database/connection"
	"github.com/alex-arraga/backend_store/internal/database/migrations"
	"github.com/alex-arraga/backend_store/pkg/auth"
	"github.com/alex-arraga/backend_store/pkg/logger"
)

type AppConfig struct {
	Port string
	DB   *gorm.DB
}

// LoadAppConfig load app and database configuration
func LoadAppConfig() (*AppConfig, error) {
	port, dbConn, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Connect to database
	db, err := connection.ConnectDatabase(dbConn)
	if err != nil {
		return nil, err
	}

	// Init other app configs
	migrations.ExecMigrations()        // Execute migrations of db
	logger.InitLogger("eccomerce_app") // Init global app logger
	auth.LoadJWTKey()                  // Load JWT Secret Key

	return &AppConfig{
		Port: port,
		DB:   db,
	}, nil
}
