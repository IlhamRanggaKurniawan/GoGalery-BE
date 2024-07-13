package entity

import "time"

type SaveContent struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ContentID uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
