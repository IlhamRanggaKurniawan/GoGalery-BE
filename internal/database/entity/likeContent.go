package entity

import "time"

type LikeContent struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ContentID uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
