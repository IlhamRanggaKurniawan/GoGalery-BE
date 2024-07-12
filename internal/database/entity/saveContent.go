package entity

import "time"

type SaveContent struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ContentID int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
