package entity

import "time"

type Feedback struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	UserId    uint64 `gorm:"index"`
	Message   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
