package entity

import "time"

type Content struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UploaderID uint
	Caption    string
	URL        string        `gorm:"unique"`
	Likes      []LikeContent `gorm:"foreignKey:ID"`
	Save       []SaveContent `gorm:"foreignKey:ID"`
	Comments   []Comment     `gorm:"foreignKey:ID"`
	Uploader   User          `gorm:"foreignKey:UploaderID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime"`
}

type LikeContent struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ContentID uint
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SaveContent struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	ContentID int
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Content   Content   `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Comment struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Comment   string
	UserID    uint
	ContentID uint
	User      User
	Content   Content
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
