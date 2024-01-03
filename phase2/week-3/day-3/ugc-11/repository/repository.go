package repository

import "gorm.io/gorm"

type Repository struct {
	User        UserInterface
	Product     ProductInterface
	Transaction TransactionInterface
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:        initUser(db),
		Product:     initProduct(db),
		Transaction: initTransaction(db),
	}
}
