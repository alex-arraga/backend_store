package repositories

import (
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	UserRepo UserRepository
}

// Connects db with the all repositories
func LoadRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		UserRepo: newUserRepo(db),
	}
}
