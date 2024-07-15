package entity

import "time"

type Notification struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Content    string
	IsChecked  bool `gorm:"default:false"`
	ReceiverID uint64
	TriggerID  uint64
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
