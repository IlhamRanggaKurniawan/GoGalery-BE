package comment

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(userId uint64, contentId uint64, text string) (*entity.Comment, error)
	FindAll(contentId uint64) (*[]entity.Comment, error)
	FindOne(id uint64) (*entity.Comment, error)
	Update(comment *entity.Comment) (*entity.Comment, error)
	DeleteOne(id uint64) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(userId uint64, contentId uint64, text string) (*entity.Comment, error) {
	comment := entity.Comment{
		UserId:    userId,
		ContentId: contentId,
		Comment:   text,
	}
	err := r.db.Create(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) FindAll(contentId uint64) (*[]entity.Comment, error) {
	var comments []entity.Comment

	err := r.db.Preload("Content").Preload("Content.Uploader").Preload("User").Where("content_id = ?", contentId).Find(&comments).Error

	if err != nil {
		return nil, err
	}

	return &comments, nil
}

func (r *commentRepository) FindOne(id uint64) (*entity.Comment, error) {
	var comment entity.Comment

	err := r.db.Where("id = ?", id).Take(&comment).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Update(comment *entity.Comment) (*entity.Comment, error) {
	r.db.Save(&comment)

	return comment, nil
}

func (r *commentRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.Comment{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
