package entity

import "time"

type Feedback struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Message   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
