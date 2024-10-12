package entity

import "time"

type User struct {
	Id                  uint64 `gorm:"primaryKey;autoIncrement"`
	Username            string `gorm:"uniqueIndex;not null"`
	Email               string `gorm:"uniqueIndex;not null"`
	Password            string `gorm:"not null"`
	Role                string `gorm:"default:member"`
	ProfileUrl          *string
	Bio                 *string
	Token               *string
	Contents            []Content        `gorm:"foreignKey:UploaderId;constraint:OnDelete:CASCADE;"`
	LikeContents        []LikeContent    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	SaveContents        []SaveContent    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Comments            []Comment        `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Followers           []Follow         `gorm:"foreignKey:FollowingId;constraint:OnDelete:CASCADE;"`
	Following           []Follow         `gorm:"foreignKey:FollowerId;constraint:OnDelete:CASCADE;"`
	SentMessages        []Message        `gorm:"foreignKey:SenderId;constraint:OnDelete:CASCADE;"`
	ParticipantInDM1    []DirectMessage  `gorm:"foreignKey:Participant1Id;constraint:OnDelete:CASCADE;"`
	ParticipantInDM2    []DirectMessage  `gorm:"foreignKey:Participant2Id;constraint:OnDelete:CASCADE;"`
	GroupChats          []GroupChat      `gorm:"many2many:group_chat_members;constraint:OnDelete:CASCADE;"`
	Notifications       []Notification   `gorm:"foreignKey:ReceiverId;constraint:OnDelete:CASCADE;"`
	NotificationTrigger []Notification   `gorm:"foreignKey:TriggerId;constraint:OnDelete:CASCADE;"`
	AIConversations     []AIConversation `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	AIMessages          []AIMessage      `gorm:"foreignKey:SenderId;constraint:OnDelete:CASCADE;"`
	Feedbacks           []Feedback       `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	CreatedAt           time.Time        `gorm:"autoCreateTime"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime"`
}
