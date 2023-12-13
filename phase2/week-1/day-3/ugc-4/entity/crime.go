package entity

import "time"

type Crime struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	HeroID      int       `json:"hero_id"`
	VillainID   int       `json:"villain_id"`
}
