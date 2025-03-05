package repositories

import "gorm.io/gorm"

type AuthRepository interface {
}

func newAuthRepo(db *gorm.DB) AuthRepository {
	return &RepoConnection{db: db}
}