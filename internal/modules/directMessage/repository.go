package directmessage

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type DirectMessageRepository interface {
	Create(participants []entity.User) (*entity.DirectMessage, error)
	FindAll(userId uint) (*[]entity.DirectMessage, error)
	FindOne(id uint) (*entity.DirectMessage, error)
	DeleteOne(id uint) error
}

type directMessageRepository struct {
	db *gorm.DB
}

func NewDirectMessageRepository(db *gorm.DB) DirectMessageRepository {
	return &directMessageRepository{db: db}
}

func (r *directMessageRepository) Create(participants []entity.User) (*entity.DirectMessage, error) {
	directMessage := entity.DirectMessage{
		Participants: participants,
	}

	err := r.db.Create(&directMessage).Error

	if err != nil {
		return nil, err
	}

	return &directMessage, nil
}

func (r *directMessageRepository) FindAll(userId uint) (*[]entity.DirectMessage, error) {

	var DirectMessage []entity.DirectMessage

	err := r.db.Joins("JOIN direct_message_participants ON direct_message_participants.direct_message_id = direct_messages.id").
		Joins("JOIN users ON users.id = direct_message_participants.user_id").
		Where("users.id = ?", userId).Preload("Participants").Find(&DirectMessage).Error

	if err != nil {
		return nil, err
	}

	return &DirectMessage, nil
}

func (r *directMessageRepository) FindOne(id uint) (*entity.DirectMessage, error) {

	var DirectMessage entity.DirectMessage

	err := r.db.Preload("Messages").Preload("Participants").Where("id = ?", id).Take(&DirectMessage).Error

	if err != nil {
		return nil, err
	}

	return &DirectMessage, nil
}

func (r *directMessageRepository) DeleteOne(id uint) error {

	err := r.db.Delete(&entity.DirectMessage{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
