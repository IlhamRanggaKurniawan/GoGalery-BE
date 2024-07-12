package entity

import "time"

type Content struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UploaderID uint
	Caption    string
	URL        string        `gorm:"unique"`
	Likes      []LikeContent `gorm:"foreignKey:ContentID"`
	Save       []SaveContent `gorm:"foreignKey:ContentID"`
	Comments   []Comment     `gorm:"foreignKey:ContentID"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}