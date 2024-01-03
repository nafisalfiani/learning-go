package repository

import (
	"avengers-commerce/entity"

	"gorm.io/gorm"
)

type product struct {
	db *gorm.DB
}

type ProductInterface interface {
	List() ([]entity.Product, error)
}

// initProduct create product repository
func initProduct(db *gorm.DB) ProductInterface {
	return &product{
		db: db,
	}
}

func (p *product) List() ([]entity.Product, error) {
	products := []entity.Product{}
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
