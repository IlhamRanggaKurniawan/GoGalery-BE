package groupchat

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type GroupChatRepository interface {
	Create(name string, members []entity.User) (*entity.GroupChat, error)
	AddMembers(groupChatID uint64, newMembers []entity.User) (*entity.GroupChat, error)
	FindAll(userId uint64) (*[]entity.GroupChat, error)
	FindOne(id uint64) (*entity.GroupChat, error)
	Update(id uint64, pictureUrl string) (*entity.GroupChat, error)
	LeaveGroup(userId uint64, groupId uint64) error
	DeleteOne(id uint64) error
}

type groupChatRepository struct {
	db *gorm.DB
}

func NewGroupChatRepository(db *gorm.DB) GroupChatRepository {
	return &groupChatRepository{db: db}
}

func (r *groupChatRepository) Create(name string, participants []entity.User) (*entity.GroupChat, error) {
	groupChat := entity.GroupChat{
		Name:    name,
		Members: participants,
	}

	err := r.db.Create(&groupChat).Error

	if err != nil {
		return nil, err
	}

	return &groupChat, nil
}

func (r *groupChatRepository) AddMembers(groupChatID uint64, newMembers []entity.User) (*entity.GroupChat, error) {
	var groupChat entity.GroupChat

	err := r.db.Preload("Members").Where("id = ?", groupChatID).Take(&groupChat).Error

	if err != nil {
		return nil, err
	}

	existingMembers := make(map[uint64]bool)
	for _, member := range groupChat.Members {
		existingMembers[member.ID] = true
	}

	for _, newMember := range newMembers {
		if !existingMembers[newMember.ID] {
			groupChat.Members = append(groupChat.Members, newMember)
		}
	}

	err = r.db.Save(&groupChat).Error

	if err != nil {
		return nil, err
	}

	return &groupChat, nil
}

func (r *groupChatRepository) FindAll(userId uint64) (*[]entity.GroupChat, error) {

	var groupChats []entity.GroupChat

	err := r.db.Joins("JOIN group_chat_members ON group_chat_members.group_chat_id = group_chats.id").
		Joins("JOIN users ON users.id = group_chat_members.user_id").
		Where("users.id = ?", userId).
		Preload("Members").
		Find(&groupChats).Error

	if err != nil {
		return nil, err
	}

	return &groupChats, nil
}

func (r *groupChatRepository) FindOne(id uint64) (*entity.GroupChat, error) {

	var groupChat entity.GroupChat

	err := r.db.Preload("Messages").Preload("Members").Where("id = ?", id).Take(&groupChat).Error

	if err != nil {
		return nil, err
	}

	return &groupChat, nil
}

func (r *groupChatRepository) Update(id uint64, pictureUrl string) (*entity.GroupChat, error) {

	groupChat, err := r.FindOne(id)

	if err != nil {
		return nil, err
	}
	groupChat.PictureUrl = &pictureUrl

	r.db.Save(&groupChat)

	return groupChat, nil
}

func (r *groupChatRepository) LeaveGroup(userId uint64, groupId uint64) error {

	var groupChat entity.GroupChat

	err := r.db.First(&groupChat, groupId).Error

	if err != nil {
		return err
	}

	err = r.db.Model(&groupChat).Association("Members").Delete(&entity.User{ID: userId})

	if err != nil {
		return err
	}

	return nil
}

func (r *groupChatRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.GroupChat{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
