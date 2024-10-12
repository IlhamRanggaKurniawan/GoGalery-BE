package entity

import "time"

type ContentType string

const (
	ImageContentType ContentType = "image"
	VideoContentType ContentType = "video"
)

type Content struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement"`
	UploaderId uint64 `gorm:"index"`
	Uploader   User   `gorm:"foreignKey:UploaderId;constraint:OnDelete:CASCADE;"`
	Caption    string
	URL        string `gorm:"unique"`
	Type       ContentType
	Likes      []LikeContent `gorm:"foreignKey:ContentId;constraint:OnDelete:CASCADE;"`
	Save       []SaveContent `gorm:"foreignKey:ContentId;constraint:OnDelete:CASCADE;"`
	Comments   []Comment     `gorm:"foreignKey:ContentId;constraint:OnDelete:CASCADE;"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}
