package directmessage

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type DirectMessageRepository interface {
	Create(participants []uint64) (*entity.DirectMessage, error)
	FindAll(userId uint64) (*[]entity.DirectMessage, error)
	FindOne(id uint64) (*entity.DirectMessage, error)
	FindOneByParticipants(participantsId []uint64) (*entity.DirectMessage, error)
	DeleteOne(id uint64) error
}

type directMessageRepository struct {
	db *gorm.DB
}

func NewDirectMessageRepository(db *gorm.DB) DirectMessageRepository {
	return &directMessageRepository{db: db}
}

func (r *directMessageRepository) Create(participants []uint64) (*entity.DirectMessage, error) {
	if len(participants) != 2 {
		return nil, fmt.Errorf("exactly two participants are required")
	}

	directMessage := entity.DirectMessage{
		Participant1ID: participants[0],
		Participant2ID: participants[1],
	}

	err := r.db.Create(&directMessage).Error

	if err != nil {
		return nil, err
	}

	return &directMessage, nil
}

func (r *directMessageRepository) FindAll(userId uint64) (*[]entity.DirectMessage, error) {

	var directMessages []entity.DirectMessage

	err := r.db.Preload("Participant1").
		Preload("Participant2").
		Where("participant1_id = ? OR participant2_id = ?", userId, userId).
		Find(&directMessages).Error

	if err != nil {
		return nil, err
	}

	return &directMessages, nil
}

func (r *directMessageRepository) FindOne(id uint64) (*entity.DirectMessage, error) {

	var DirectMessage entity.DirectMessage

	err := r.db.Preload("Messages").Preload("Participant1").Preload("Participant2").Where("id = ?", id).Take(&DirectMessage).Error

	if err != nil {
		return nil, err
	}

	return &DirectMessage, nil
}

func (r *directMessageRepository) FindOneByParticipants(participantsId []uint64) (*entity.DirectMessage, error) {
	if len(participantsId) != 2 {
		return nil, fmt.Errorf("exactly two participants are required")
	}

	var directMessage entity.DirectMessage

	err := r.db.Preload("Participant1").
		Preload("Participant2").
		Where("(participant1_id = ? AND participant2_id = ?) OR (participant1_id = ? AND participant2_id = ?)",
			participantsId[0], participantsId[1], participantsId[1], participantsId[0]).
		First(&directMessage).Error

	if err != nil {
		return nil, err
	}

	return &directMessage, nil
}

func (r *directMessageRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.DirectMessage{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
