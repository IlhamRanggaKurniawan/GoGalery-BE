package like

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type LikeContentRepository interface {
	Create(userId uint64, contentId uint64) (*entity.LikeContent, error)
	FindOne(userId uint64, contentId uint64) (*entity.LikeContent, error)
	DeleteOne(id uint64) error
}

type likeContentRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeContentRepository {
	return &likeContentRepository{db: db}
}

func (r *likeContentRepository) Create(userId uint64, contentId uint64) (*entity.LikeContent, error) {
	like := entity.LikeContent{
		UserID:    userId,
		ContentID: contentId,
	}

	err := r.db.Create(&like).Error

	if err != nil {
		return nil, err
	}

	return &like, nil
}

func (r *likeContentRepository) FindOne(userId uint64, contentId uint64) (*entity.LikeContent, error) {
	var like entity.LikeContent

	err := r.db.Where("user_id = ? AND content_id = ?", userId, contentId).Take(&like).Error

	if err != nil {
		return nil, err
	}

	return &like, nil
}

func (r *likeContentRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.LikeContent{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
