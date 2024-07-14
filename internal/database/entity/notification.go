package entity

import "time"

type Notification struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Content    string
	IsChecked  bool `gorm:"default:false"`
	ReceiverID uint
	TriggerID  uint
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
