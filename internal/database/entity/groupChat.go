package entity

import "time"

type GroupChat struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	PictureUrl *string
	Members    []User    `gorm:"many2many:group_chat_members;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
