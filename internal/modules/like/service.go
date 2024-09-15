package like

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type LikeContentService interface {
	LikeContent(userId uint64, contentId uint64) (*entity.LikeContent, error)
	GetOneLike(userId uint64, contentId uint64) (*entity.LikeContent, error)
	UnlikeContent(id uint64) error
}

type likeContentService struct {
	likeContentRepository LikeContentRepository
}

func NewLikeService(likeContentRepository LikeContentRepository) LikeContentService {
	return &likeContentService{
		likeContentRepository: likeContentRepository,
	}
}

func (s *likeContentService) LikeContent(userId uint64, contentId uint64) (*entity.LikeContent, error) {

	like, err := s.likeContentRepository.Create(userId, contentId)

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (s *likeContentService) GetOneLike(userId uint64, contentId uint64) (*entity.LikeContent, error) {

	like, err := s.likeContentRepository.FindOne(userId, contentId)

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (s *likeContentService) UnlikeContent(id uint64) error {

	err := s.likeContentRepository.DeleteOne(id)

	return err
}
