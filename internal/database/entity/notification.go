package entity

import "time"

type Notification struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Content    string
	Type       string
	IsRead     bool `gorm:"default:false"`
	ReceiverID uint
	TriggerID  uint
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
