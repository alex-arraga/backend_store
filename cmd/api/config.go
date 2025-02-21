package main

import (
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/config"
	"github.com/alex-arraga/backend_store/internal/database/connection"
	"github.com/alex-arraga/backend_store/internal/database/migrations"
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

	// Execute migrations of db
	migrations.ExecMigrations()

	return &AppConfig{
		Port: port,
		DB:   db,
	}, nil
}
