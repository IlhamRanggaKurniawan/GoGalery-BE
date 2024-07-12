package entity

import "time"

type AIMessage struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Message        string
	Response       *string
	SenderID       uint
	ConversationID uint
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}


