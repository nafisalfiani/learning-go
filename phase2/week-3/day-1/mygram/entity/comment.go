package entity

import (
	"time"
)

// Comment entity
type Comment struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
