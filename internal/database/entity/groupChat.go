package entity

import "time"

type GroupChat struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"index"`
	PictureUrl *string
	Members    []User    `gorm:"many2many:group_chat_members;constraint:OnDelete:CASCADE;"`
	Messages   []Message `gorm:"foreignKey:GroupChatId;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
