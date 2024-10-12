package entity

import "time"

type Notification struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement"`
	Content    string
	IsChecked  bool      `gorm:"default:false"`
	ReceiverId uint64    `gorm:"index"`
	TriggerId  uint64    `gorm:"index"`
	Receiver   User      `gorm:"foreignKey:ReceiverId;constraint:OnDelete:CASCADE;"`
	Trigger    User      `gorm:"foreignKey:TriggerId;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
