package entity

import "time"

type Comment struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Comment   string
	UserID    uint64
	ContentID uint64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
