package entity

import "time"

type Comment struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	Comment   string
	UserId    uint64    `gorm:"index"`
	ContentId uint64    `gorm:"index"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Content   Content   `gorm:"foreignKey:ContentId;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
