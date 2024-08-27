package entity

import "time"

type SaveContent struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	UserID    uint64
	ContentID uint64
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
