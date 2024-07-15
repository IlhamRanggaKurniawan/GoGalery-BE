package feedback

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(userId uint64, message string) (*entity.Feedback, error)
	FindAll() (*[]entity.Feedback, error)
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) Create(userId uint64, message string) (*entity.Feedback, error) {
	feedback := entity.Feedback{
		UserID:  userId,
		Message: message,
	}

	err := r.db.Create(&feedback).Error

	if err != nil {
		return nil, err
	}

	return &feedback, nil
}

func (r *feedbackRepository) FindAll() (*[]entity.Feedback, error) {
	var feedbacks []entity.Feedback

	err := r.db.Find(&feedbacks).Error

	if err != nil {
		return nil, err
	}

	return &feedbacks, nil
}
