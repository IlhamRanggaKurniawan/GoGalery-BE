package entity

import "time"

type DirectMessage struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement"`
	Participant1ID uint64    `gorm:"index:idx_dm_participant"`
	Participant2ID uint64    `gorm:"index:idx_dm_participant"`
	Participant1   User      `gorm:"foreignKey:Participant1ID"`
	Participant2   User      `gorm:"foreignKey:Participant2ID"`
	Messages       []Message `gorm:"foreignKey:DirectMessageID;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
