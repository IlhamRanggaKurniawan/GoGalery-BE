package save

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type SaveContentRepository interface {
	Create(userId uint, contentId uint) (*entity.SaveContent, error)
	FindAll(contentId uint) (*[]entity.SaveContent, error)
	FindOne(userId uint, contentId uint) (*entity.SaveContent, error)
	DeleteOne(id uint) error
}

type saveContentRepository struct {
	db *gorm.DB
}

func NewSaveRepository(db *gorm.DB) SaveContentRepository {
	return &saveContentRepository{db: db}
}

func (r *saveContentRepository) Create(userId uint, contentId uint) (*entity.SaveContent, error) {
	save := entity.SaveContent{
		UserID:    userId,
		ContentID: contentId,
	}

	err := r.db.Create(&save).Error

	if err != nil {
		return nil, err
	}

	return &save, nil
}

func (r *saveContentRepository) FindAll(contentId uint) (*[]entity.SaveContent, error) {
	var saves []entity.SaveContent

	err := r.db.Where("content_id = ?", contentId).Find(&saves).Error

	if err != nil {
		return nil, err
	}

	return &saves, nil
}

func (r *saveContentRepository) FindOne(userId uint, contentId uint) (*entity.SaveContent, error) {
	var save entity.SaveContent

	err := r.db.Where("user_id = ? AND content_id = ?", userId, contentId).Take(&save).Error

	if err != nil {
		return nil, err
	}

	return &save, nil
}

func (r *saveContentRepository) DeleteOne(id uint) error {

	err := r.db.Delete(&entity.SaveContent{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
