package entity

import "time"

type AIMessage struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Message        string
	Response       *string
	SenderID       uint
	ConversationID uint
	Sender         User           `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Conversation   AIConversation `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
}

type AIConversation struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
