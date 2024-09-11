package entity

import "time"

type AIMessage struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Message        string
	Response       *string
	SenderID       uint64    `gorm:"index"`
	ConversationID uint64    `gorm:"index"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
