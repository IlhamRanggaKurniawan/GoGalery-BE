package entity

import "time"

type LikeContent struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"index"`
	ContentID uint64    `gorm:"index"`
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
