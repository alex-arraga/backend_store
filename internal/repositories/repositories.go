package repositories

import (
	"gorm.io/gorm"
)

type RepoConnection struct {
	db *gorm.DB
}
type RepositoryContainer struct {
	AuthAccountRepo AuthAccountRepository
	UserRepo        UserRepository
}

// Connects db with the all repositories
func LoadRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		AuthAccountRepo: newAuthAccountRepo(db),
		UserRepo:        newUserRepo(db),
	}
}
