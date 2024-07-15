package entity

import "time"

type AIMessage struct {
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Message        string
	Response       *string
	SenderID       uint64
	ConversationID uint64
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}


