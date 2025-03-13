package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type AuthAccountRepository interface {
	someFunc() error
}

func newAuthAccountRepo(db *gorm.DB) AuthAccountRepository {
	return &RepoConnection{db: db}
}

func (repo *RepoConnection) someFunc() error {
	return errors.New("some error")
}
