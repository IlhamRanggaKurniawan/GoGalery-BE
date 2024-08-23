package entity

import "time"

type ContentType string

const (
	ImageContentType ContentType = "image"
	VideoContentType ContentType = "video"
)

type Content struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	UploaderID uint64
	Uploader   User `gorm:"foreignKey:UploaderID;constraint:OnDelete:CASCADE;"`
	Caption    string
	URL        string `gorm:"unique"`
	Type       ContentType
	Likes      []LikeContent `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"`
	Save       []SaveContent `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"`
	Comments   []Comment     `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}
