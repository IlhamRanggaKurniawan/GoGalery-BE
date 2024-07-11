package entity

import "time"

type User struct {
	ID                  uint   `gorm:"primaryKey;autoIncrement"`
	Username            string `gorm:"unique"`
	Email               string `gorm:"unique"`
	Password            string
	Role                string `gorm:"default:member"`
	ProfileUrl          *string
	Bio                 *string
	Contents            []Content        `gorm:"foreignKey:UploaderID"`
	LikeContents        []LikeContent    `gorm:"foreignKey:UserID"`
	SaveContents        []SaveContent    `gorm:"foreignKey:UserID"`
	Comments            []Comment        `gorm:"foreignKey:UserID"`
	Followers           []Follow         `gorm:"foreignKey:FollowingID"`
	Following           []Follow         `gorm:"foreignKey:FollowerID"`
	Messages            []Message        `gorm:"foreignKey:SenderID"`
	DirectMessages      []DirectMessage  `gorm:"many2many:user_direct_messages;"`
	GroupChats          []GroupChat      `gorm:"many2many:user_group_chats;"`
	Notifications       []Notification   `gorm:"foreignKey:ReceiverID"`
	NotificationTrigger []Notification   `gorm:"foreignKey:TriggerID"`
	AIConversation      []AIConversation `gorm:"foreignKey:UserID"`
	AIMessages          []AIMessage      `gorm:"foreignKey:SenderID"`
	Feedback            []Feedback       `gorm:"foreignKey:UserID"`
	CreatedAt           time.Time        `gorm:"autoCreateTime"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime"`
}

type Follow struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	FollowerID  uint
	FollowingID uint
	Follower    User      `gorm:"foreignKey:FollowerID;references:ID"`
	Following   User      `gorm:"foreignKey:FollowingID;references:ID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type Notification struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Content    string
	Type       string
	IsRead     bool
	ReceiverID uint
	TriggerID  uint
	Receiver   User      `gorm:"foreignKey:ReceiverID;references:ID"`
	Trigger    User      `gorm:"foreignKey:TriggerID;references:ID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
