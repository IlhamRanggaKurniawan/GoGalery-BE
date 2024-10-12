package entity

import "time"

type Message struct {
	Id              uint64 `gorm:"primaryKey;autoIncrement"`
	Message         string
	IsRead          bool      `gorm:"default:false"`
	SenderId        uint64    `gorm:"index"`
	DirectMessageId *uint64   `gorm:"index"`
	GroupChatId     *uint64   `gorm:"index"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
