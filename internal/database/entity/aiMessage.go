package entity

import "time"

type AIMessage struct {
	Id             uint64 `gorm:"primaryKey;autoIncrement"`
	Message        string
	Response       *string
	SenderId       uint64    `gorm:"index"`
	ConversationId uint64    `gorm:"index"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
