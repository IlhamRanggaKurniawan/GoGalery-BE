package entity

import "time"

type SaveContent struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	UserId    uint64    `gorm:"index"`
	ContentId uint64    `gorm:"index"`
	Content   Content   `gorm:"foreignKey:ContentId;constraint:OnDelete:CASCADE;"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
