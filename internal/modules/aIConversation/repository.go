package aIconversation

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type AIConversationRepository interface {
	Create(userId uint64) (*entity.AIConversation, error)
	FindOne(userId uint64) (*entity.AIConversation, error)
	DeleteOne(id uint64) error
}

type aIConversationRepository struct {
	db *gorm.DB
}

func NewAIConversationRepository(db *gorm.DB) AIConversationRepository {
	return &aIConversationRepository{db: db}
}

func (r *aIConversationRepository) Create(userId uint64) (*entity.AIConversation, error) {
	message := entity.AIConversation{
		UserId: userId,
	}

	err := r.db.Create(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *aIConversationRepository) FindOne(userId uint64) (*entity.AIConversation, error) {

	var aIConversation entity.AIConversation

	err := r.db.Preload("Messages").Where("user_id = ?", userId).Take(&aIConversation).Error

	if err != nil {
		return nil, err
	}

	return &aIConversation, nil
}

func (r *aIConversationRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.AIConversation{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
