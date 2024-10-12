package entity

import "time"

type Follow struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement"`
	FollowerId  uint64    `gorm:"index:idx_follow_relation"`
	FollowingId uint64    `gorm:"index:idx_follow_relation"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
