package entity

import "time"

type DirectMessage struct {
	ID           uint64      `gorm:"primaryKey;autoIncrement"`
	Participants []User    `gorm:"many2many:direct_message_participants;constraint:OnDelete:CASCADE;"`
	Messages     []Message `gorm:"foreignKey:DirectMessageID;constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
