package aImessage

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type AIMessageRepository interface {
	Create(senderId uint64, conversationId uint64, message string, response string) (*entity.AIMessage, error)
	FindAll(conversationId uint64) (*[]entity.AIMessage, error)
	FindOne(id uint64) (*entity.AIMessage, error)
	Update(id uint64, message string) (*entity.AIMessage, error)
	DeleteOne(id uint64) error
}

type aIMessageRepository struct {
	db *gorm.DB
}

func NewAIMessageRepository(db *gorm.DB) AIMessageRepository {
	return &aIMessageRepository{db: db}
}

func (r *aIMessageRepository) Create(senderId uint64, conversationId uint64, message string, response string) (*entity.AIMessage, error) {
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

func (r *aIMessageRepository) FindAll(conversationId uint64) (*[]entity.AIMessage, error) {
	var aiMessages []entity.AIMessage

	err := r.db.Where("conversation_id = ?", conversationId).Find(&aiMessages).Error

	if err != nil {
		return nil, err
	}

	return &aiMessages, nil
}

func (r *aIMessageRepository) FindOne(id uint64) (*entity.AIMessage, error) {
	var aIMessage entity.AIMessage

	err := r.db.Where("id = ?", id).Take(&aIMessage).Error

	if err != nil {
		return nil, err
	}

	return &aIMessage, nil
}

func (r *aIMessageRepository) Update(id uint64, message string) (*entity.AIMessage, error) {
	aIMessage, err := r.FindOne(id)

	if err != nil {
		return nil, err
	}
	aIMessage.Message = message

	r.db.Save(&message)

	return aIMessage, nil
}

func (r *aIMessageRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.AIMessage{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
