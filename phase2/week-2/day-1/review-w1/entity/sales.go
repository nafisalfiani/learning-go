package entity

import "time"

type Sales struct {
	Id       int       `json:"id"`
	GameId   int       `json:"game_id"`
	BranchId int       `json:"branch_id"`
	Date     time.Time `json:"date"`
	Quantity int       `json:"quantity"`
}
