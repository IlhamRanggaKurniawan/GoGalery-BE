package entity

import "time"

type Feedback struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Message   string
	User      User      `gorm:"foreignKey:UserID;reference:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
