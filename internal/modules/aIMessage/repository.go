package aImessage

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type AIMessageRepository interface {
	Create(senderId uint64, conversationId uint64, message string, response string) (*entity.AIMessage, error)
	FindOne(id uint64) (*entity.AIMessage, error)
	Update(message *entity.AIMessage) (*entity.AIMessage, error)
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
		SenderId:       senderId,
		ConversationId: conversationId,
		Message:        message,
		Response:       &response,
	}

	err := r.db.Create(&aIMessage).Error

	if err != nil {
		return nil, err
	}

	return &aIMessage, nil
}

func (r *aIMessageRepository) FindOne(id uint64) (*entity.AIMessage, error) {
	var message entity.AIMessage

	err := r.db.Where("id = ?", id).Take(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *aIMessageRepository) Update(message *entity.AIMessage) (*entity.AIMessage, error) {

	r.db.Save(&message)

	return message, nil
}

func (r *aIMessageRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.AIMessage{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
