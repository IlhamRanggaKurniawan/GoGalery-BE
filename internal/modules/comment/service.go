package comment

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type CommentService interface {
	SendComment(userId uint, contentId uint, text string) (*entity.Comment, error)
	updateComment(id uint, text string) (*entity.Comment, error)
	GetAllComments(contentId uint) (*[]entity.Comment, error)
	DeleteContent(id uint) error
}

type commentService struct {
	commentRepository CommentRepository
}

func NewContentService(commentRepository CommentRepository) CommentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) SendComment(userId uint, contentId uint, text string) (*entity.Comment, error) {
	content, err := s.commentRepository.Create(userId, contentId, text)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *commentService) updateComment(id uint, text string) (*entity.Comment, error) {
	content, err := s.commentRepository.Update(id, text)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *commentService) GetAllComments(contentId uint) (*[]entity.Comment, error) {
	content, err := s.commentRepository.FindAll(contentId)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *commentService) DeleteContent(id uint) error {
	err := s.commentRepository.DeleteOne(id)

	return err
}
