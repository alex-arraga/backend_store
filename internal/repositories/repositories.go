package repositories

import (
	"gorm.io/gorm"
)

type RepoConnection struct {
	db *gorm.DB
}
type RepositoryContainer struct {
	UserRepo UserRepository
	AuthAccountRepo AuthAccountRepository
}

// Connects db with the all repositories
func LoadRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		UserRepo: newUserRepo(db),
		AuthAccountRepo: newAuthAccountRepo(db),
	}
}
