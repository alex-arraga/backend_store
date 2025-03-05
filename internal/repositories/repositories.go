package repositories

import (
	"gorm.io/gorm"
)

type RepoConnection struct {
	db *gorm.DB
}
type RepositoryContainer struct {
	AuthRepo AuthRepository
	UserRepo UserRepository
}

// Connects db with the all repositories
func LoadRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		AuthRepo: newAuthRepo(db),
		UserRepo: newUserRepo(db),
	}
}
