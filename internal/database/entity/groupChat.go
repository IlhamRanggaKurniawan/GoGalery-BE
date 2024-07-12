package entity

import "time"

type GroupChat struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	PictureUrl *string
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
