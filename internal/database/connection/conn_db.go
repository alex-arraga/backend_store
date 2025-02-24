package connection

import (
	"time"

	"github.com/alex-arraga/backend_store/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase(connStr string) (*gorm.DB, error) {
	// Connection to db using GORM
	conn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Info),
	})
	if err != nil {
		logger.UseLogger().Fatal().Err(err).Msg("Error connecting database")
	}

	// Low level connection
	sqlDB, err := conn.DB()
	if err != nil {
		logger.UseLogger().Fatal().Err(err).Msg("Error returning sql DB pointer")
	}

	// Pool connections config
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	logger.UseLogger().Info().Msg("Database connection established")
	DB = conn

	return DB, nil
}
