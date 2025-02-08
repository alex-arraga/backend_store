package migrations

import (
	"log"

	"github.com/alex-arraga/backend_store/database/connection"
	"github.com/alex-arraga/backend_store/database/gorm_models"
)

func migrate(models ...interface{}) {
	for _, model := range models {
		if err := connection.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Error executing migration: %v", err)
		}
		log.Printf("Migration '%v' successfully", model)
	}
}

func ExecMigrations() {
		migrate(&gorm_models.User{})
}