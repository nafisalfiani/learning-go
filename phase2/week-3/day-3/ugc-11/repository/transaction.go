package repository

import (
	"avengers-commerce/entity"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

type TransactionInterface interface {
	List(userId int) ([]entity.Transaction, error)
	Create(newTransaction entity.Transaction) (entity.Transaction, error)
}

// initTransaction create transaction repository
func initTransaction(db *gorm.DB) TransactionInterface {
	return &transaction{
		db: db,
	}
}

func (t *transaction) List(userId int) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}
	if err := t.db.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transaction) Create(newTransaction entity.Transaction) (entity.Transaction, error) {
	if err := t.db.Transaction(func(tx *gorm.DB) error {
		product := entity.Product{ProductID: newTransaction.ProductID}
		if err := tx.First(&product).Error; err != nil {
			return err
		}

		product.ProductID = product.ProductID - newTransaction.Quantity
		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		user := entity.User{UserID: newTransaction.UserID}
		if err := tx.First(&user).Error; err != nil {
			return err
		}

		newTransaction.TotalAmount = product.Price * float64(newTransaction.Quantity)
		user.DepositAmount = user.DepositAmount - newTransaction.TotalAmount
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(&newTransaction).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
