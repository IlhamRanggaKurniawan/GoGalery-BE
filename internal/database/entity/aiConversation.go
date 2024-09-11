package entity

import "time"

type AIConversation struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement"`
	UserID    uint64      `gorm:"index"`
	Messages  []AIMessage `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}
