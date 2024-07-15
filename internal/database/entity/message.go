package entity

import "time"

type Message struct {
    ID              uint64      `gorm:"primaryKey;autoIncrement"`
    Message         string
    IsRead          bool      `gorm:"default:false"`
    SenderID        uint64
    DirectMessageID *uint64
    GroupChatID     *uint64
    CreatedAt       time.Time `gorm:"autoCreateTime"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
