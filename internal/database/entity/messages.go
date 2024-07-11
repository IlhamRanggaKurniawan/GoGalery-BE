package entity

import "time"

type Message struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Message         string
	IsRead          bool
	SenderID        uint
	DirectMessageID *uint
	GroupChatID     *uint
	Sender          User           `gorm:"foreignKey:SenderID"`
	DirectMessage   *DirectMessage `gorm:"foreignKey:DirectMessageID"`
	GroupChat       *GroupChat     `gorm:"foreignKey:GroupChatID"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
}

type DirectMessage struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Participants []User    `gorm:"many2many:direct_message_participants"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type GroupChat struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	PictureUrl *string
	Member     []User    `gorm:"many2many:group_chat_members"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
