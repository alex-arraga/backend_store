package migrations

import (
	"log"

	"github.com/alex-arraga/backend_store/database/connection"
)

func Migrate(models ...interface{}) {
	for _, model := range models {
		if err := connection.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Error executing migration: %v", err)
		}
		log.Printf("Migration '%v' successfully", model)
	}
}
