package entity

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"item_name"`
	Code        string `json:"item_code"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
