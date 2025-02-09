package migrations

import (
	"log"

	"github.com/alex-arraga/backend_store/database/connection"
	"github.com/alex-arraga/backend_store/database/gorm_models"
)

func migrate(models ...interface{}) {
	for _, model := range models {
		if err := connection.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Error executing migration: %d", err)
		}
	}
	log.Printf("Migrations successfully")
}

func ExecMigrations() {
	migrate(&gorm_models.User{})
}
