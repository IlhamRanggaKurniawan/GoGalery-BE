package entity

import "time"

type Message struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Message         string
	IsRead          bool
	SenderID        uint
	DirectMessageID *uint
	GroupChatID     *uint
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
}
