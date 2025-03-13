package migrations

import (
	"github.com/alex-arraga/backend_store/internal/database/connection"
	"github.com/alex-arraga/backend_store/internal/database/gorm_models"
	"github.com/alex-arraga/backend_store/pkg/logger"
)

func migrate(models ...interface{}) {
	for _, model := range models {
		if err := connection.DB.AutoMigrate(model); err != nil {
			logger.UseLogger().Fatal().Err(err).Str("module", "database").Msg("Error executing migrations")
		}
	}
	logger.UseLogger().Debug().Str("module", "database").Msg("Migrations successfully executed")
}

func ExecMigrations() {
	migrate(
		&gorm_models.User{},
		&gorm_models.AuthAccount{},
	)
}

// To Devs: Execute this function inside ExecMigrations to delete a table
// func resetTables() {
// 	table := &gorm_models.User{}
// 	connection.DB.Migrator().DropTable(table)
// }
