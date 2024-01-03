package entity

type TransactionRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type Transaction struct {
	TransactionID int     `json:"transaction_id" gorm:"primaryKey"`
	UserID        int     `json:"user_id"`
	ProductID     int     `json:"product_id"`
	Quantity      int     `json:"quantity"`
	TotalAmount   float64 `json:"total_amount"`
}
