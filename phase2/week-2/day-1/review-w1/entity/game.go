package entity

type Game struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
