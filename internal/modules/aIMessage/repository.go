package aImessage

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type AIMessageRepository interface {
	Create(senderId uint, conversationId uint, message string, response string) (*entity.AIMessage, error)
	FindAll(conversationId uint) (*[]entity.AIMessage, error)
	FindOne(id uint) (*entity.AIMessage, error)
	Update(id uint, message string) (*entity.AIMessage, error)
	DeleteOne(id uint) error
}

type aIMessageRepository struct {
	db *gorm.DB
}

func NewAIMessageRepository(db *gorm.DB) AIMessageRepository {
	return &aIMessageRepository{db: db}
}

func (r *aIMessageRepository) Create(senderId uint, conversationId uint, message string, response string) (*entity.AIMessage, error) {
	aIMessage := entity.AIMessage{
		SenderID: senderId,
		ConversationID: conversationId,
		Message:  message,
		Response: &response,
	}

	err := r.db.Create(&aIMessage).Error

	if err != nil {
		return nil, err
	}

	return &aIMessage, nil
}

func (r *aIMessageRepository) FindAll(conversationId uint) (*[]entity.AIMessage, error) {
	var aiMessages []entity.AIMessage

	err := r.db.Where("conversation_id = ?", conversationId).Find(&aiMessages).Error

	if err != nil {
		return nil, err
	}

	return &aiMessages, nil
}

func (r *aIMessageRepository) FindOne(id uint) (*entity.AIMessage, error) {
	var aIMessage entity.AIMessage

	err := r.db.Where("id = ?", id).Take(&aIMessage).Error

	if err != nil {
		return nil, err
	}

	return &aIMessage, nil
}

func (r *aIMessageRepository) Update(id uint, message string) (*entity.AIMessage, error) {
	aIMessage, err := r.FindOne(id)

	if err != nil {
		return nil, err
	}
	aIMessage.Message = message

	r.db.Save(&message)

	return aIMessage, nil
}

func (r *aIMessageRepository) DeleteOne(id uint) error {

	err := r.db.Delete(&entity.AIMessage{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
