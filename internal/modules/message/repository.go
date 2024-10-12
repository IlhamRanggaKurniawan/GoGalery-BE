package message

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(senderId uint64, directMessageId uint64, groupChatId uint64, text string) (*entity.Message, error)
	FindAll(directMessageId uint64, groupChatId uint64) (*[]entity.Message, error)
	FindOne(id uint64) (*entity.Message, error)
	Update(id uint64, text string) (*entity.Message, error)
	DeleteOne(id uint64) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(senderId uint64, directMessageId uint64, groupChatId uint64, text string) (*entity.Message, error) {
	var message entity.Message
	if directMessageId != 0 {
		message = entity.Message{
			Message:         text,
			SenderId:        senderId,
			DirectMessageId: &directMessageId,
			GroupChatId:     nil,
		}
	} else {
		message = entity.Message{
			Message:         text,
			SenderId:        senderId,
			DirectMessageId: nil,
			GroupChatId:     &groupChatId,
		}
	}

	err := r.db.Create(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *messageRepository) FindAll(directMessageId uint64, groupChatId uint64) (*[]entity.Message, error) {
	var messages []entity.Message

	err := r.db.Where("direct_message_id = ? OR group_chat_id = ?", directMessageId, groupChatId).Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return &messages, nil
}

func (r *messageRepository) FindOne(id uint64) (*entity.Message, error) {
	var message entity.Message

	err := r.db.Where("id = ?", id).Take(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *messageRepository) Update(id uint64, text string) (*entity.Message, error) {
	message, err := r.FindOne(id)

	if err != nil {
		return nil, err
	}
	message.Message = text

	r.db.Save(&message)

	return message, nil
}

func (r *messageRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.Message{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
