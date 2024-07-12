package entity

import "time"

type Comment struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Comment   string
	UserID    uint
	ContentID uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
