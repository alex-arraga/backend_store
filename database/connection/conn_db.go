package connection

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase(connStr string) (*gorm.DB, error) {
	// Connection to db using GORM
	conn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting database: %d", err)
	}

	// Low level connection
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("Error returning DB pointer: %d", err)
	}

	// Pool connections config
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Print("Database connection established")
	DB = conn

	return DB, nil
}
