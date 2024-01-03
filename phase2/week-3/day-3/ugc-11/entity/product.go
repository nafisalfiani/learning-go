package entity

type Product struct {
	ProductID int     `json:"product_id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}
