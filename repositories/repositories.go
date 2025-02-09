package repositories

import "gorm.io/gorm"

// Connects db with the all repositories
func LoadRepositories(db *gorm.DB) {
	newUserRepo(db)
}