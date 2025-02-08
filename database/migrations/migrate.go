package migrations

import (
	"log"

	"github.com/alex-arraga/backend_store/database/connection"
)

func Migrate(models ...interface{}) {
	err := connection.DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Error executing migration: %d", err)
	}

	log.Printf("Migration '%v' successfully", models)
}
