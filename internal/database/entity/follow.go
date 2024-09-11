package entity

import "time"

type Follow struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	FollowerID  uint64    `gorm:"index:idx_follow_relation"`
	FollowingID uint64    `gorm:"index:idx_follow_relation"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
