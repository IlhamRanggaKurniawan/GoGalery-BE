package entity

import "time"

type Message struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Message         string
	IsRead          bool      `gorm:"default:false"`
	SenderID        uint64    `gorm:"index"`
	DirectMessageID *uint64   `gorm:"index"`
	GroupChatID     *uint64   `gorm:"index"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
