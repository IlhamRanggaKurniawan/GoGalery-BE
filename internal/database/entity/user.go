package entity

import "time"

type User struct {
	ID                  uint64 `gorm:"primaryKey;autoIncrement"`
	Username            string `gorm:"unique"`
	Email               string `gorm:"unique"`
	Password            string
	Role                string `gorm:"default:member"`
	ProfileUrl          *string
	Bio                 *string
	Token               *string
	Contents            []Content        `gorm:"foreignKey:UploaderID;constraint:OnDelete:CASCADE;"`
	LikeContents        []LikeContent    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	SaveContents        []SaveContent    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Comments            []Comment        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Followers           []Follow         `gorm:"foreignKey:FollowingID;constraint:OnDelete:CASCADE;"`
	Following           []Follow         `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE;"`
	Messages            []Message        `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE;"`
	DirectMessages      []DirectMessage  `gorm:"many2many:direct_message_participants;constraint:OnDelete:CASCADE;"`
	GroupChats          []GroupChat      `gorm:"many2many:group_chat_members;constraint:OnDelete:CASCADE;"`
	Notifications       []Notification   `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE;"`
	NotificationTrigger []Notification   `gorm:"foreignKey:TriggerID;constraint:OnDelete:CASCADE;"`
	AIConversation      []AIConversation `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	AIMessages          []AIMessage      `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE;"`
	Feedback            []Feedback       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt           time.Time        `gorm:"autoCreateTime"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime"`
}
