package main

import (
	"gorm.io/gorm"

	"github.com/alex-arraga/backend_store/config"
	"github.com/alex-arraga/backend_store/database/connection"
)

type AppConfig struct {
	Port string
	DB   *gorm.DB
}

// LoadAppConfig carga la configuración y la conexión a la DB.
func LoadAppConfig() (*AppConfig, error) {
	port, dbConn, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Conectar a la base de datos
	db, err := connection.ConnectDatabase(dbConn)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Port: port,
		DB:   db,
	}, nil
}
