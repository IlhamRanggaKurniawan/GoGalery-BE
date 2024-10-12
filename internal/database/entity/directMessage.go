package entity

import "time"

type DirectMessage struct {
	Id             uint64    `gorm:"primaryKey;autoIncrement"`
	Participant1Id uint64    `gorm:"index:idx_dm_participant"`
	Participant2Id uint64    `gorm:"index:idx_dm_participant"`
	Participant1   User      `gorm:"foreignKey:Participant1Id"`
	Participant2   User      `gorm:"foreignKey:Participant2Id"`
	Messages       []Message `gorm:"foreignKey:DirectMessageId;constraint:OnDelete:CASCADE;"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
