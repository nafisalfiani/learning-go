package entity

import (
	"time"
)

type Photo struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Caption   string
	PhotoURL  string
	UserID    int
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
}
