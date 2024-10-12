package entity

import "time"

type AIConversation struct {
	Id        uint64      `gorm:"primaryKey;autoIncrement"`
	UserId    uint64      `gorm:"index"`
	Messages  []AIMessage `gorm:"foreignKey:ConversationId;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}
