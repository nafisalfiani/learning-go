package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User UserInterface
}

// Init create new Repository object
func Init(db *gorm.DB) *Repository {
	return &Repository{
		User: initUser(db),
	}
}
