package entity

import "time"

type SaveContent struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	UserID    uint64
	ContentID uint64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
