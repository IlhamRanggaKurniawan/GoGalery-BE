package entity

import "time"

type Follow struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	FollowerID  uint64
	FollowingID uint64
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}