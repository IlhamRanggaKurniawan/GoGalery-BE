package entity

import "time"

type DirectMessage struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Participants []User    `gorm:"many2many:direct_message_participants"`
	Messages     []Message `gorm:"foreignKey:DirectMessageID"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
