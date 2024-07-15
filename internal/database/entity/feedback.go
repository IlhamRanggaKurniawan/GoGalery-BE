package entity

import "time"

type Feedback struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	UserID    uint64
	Message   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
