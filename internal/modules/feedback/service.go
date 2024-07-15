package feedback

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type FeedbackService interface {
	SendFeedback(userId uint64, message string) (*entity.Feedback, error)
	GetAllFeedbacks() (*[]entity.Feedback, error)
}

type feedbackService struct {
	feedbackRepository FeedbackRepository
}

func NewFeedbackService(feedbackRepository FeedbackRepository) FeedbackService {
	return &feedbackService{
		feedbackRepository: feedbackRepository,
	}
}

func (s *feedbackService) SendFeedback(userId uint64, message string) (*entity.Feedback, error) {

	feedback, err := s.feedbackRepository.Create(userId, message)

	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (s *feedbackService) GetAllFeedbacks() (*[]entity.Feedback, error) {

	feedbacks, err := s.feedbackRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}
