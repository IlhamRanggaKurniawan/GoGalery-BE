package entity

import "time"

type Follow struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	FollowerID  uint
	FollowingID uint
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}