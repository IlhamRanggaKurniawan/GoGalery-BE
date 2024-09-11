package entity

import "time"

type Notification struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Content    string
	IsChecked  bool      `gorm:"default:false"`
	ReceiverID uint64    `gorm:"index"`
	TriggerID  uint64    `gorm:"index"`
	Receiver   User      `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE;"`
	Trigger    User      `gorm:"foreignKey:TriggerID;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
