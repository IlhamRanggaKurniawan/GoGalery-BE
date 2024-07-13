package like

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type LikeContentService interface {
	LikeContent(userId uint, contentId uint) (*entity.LikeContent, error)
	GetAllLikes(contentId uint) (*[]entity.LikeContent, error)
	GetOneLike(userId uint, contentId uint) (*entity.LikeContent, error)
	UnlikeContent(id uint) error
}

type likeContentService struct {
	likeContentRepository LikeContentRepository
}

func NewLikeService(likeContentRepository LikeContentRepository) LikeContentService {
	return &likeContentService{
		likeContentRepository: likeContentRepository,
	}
}

func (s *likeContentService) LikeContent(userId uint, contentId uint) (*entity.LikeContent, error) {

	like, err := s.likeContentRepository.Create(userId, contentId)

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (s *likeContentService) GetAllLikes(contentId uint) (*[]entity.LikeContent, error) {

	likes, err := s.likeContentRepository.FindAll(contentId)

	if err != nil {
		return nil, err
	}

	return likes, nil
}

func (s *likeContentService) GetOneLike(userId uint, contentId uint) (*entity.LikeContent, error) {

	like, err := s.likeContentRepository.FindOne(userId, contentId)

	if err != nil {
		return nil, err
	}

	return like, nil
}

func (s *likeContentService) UnlikeContent(id uint) error {

	err := s.likeContentRepository.DeleteOne(id)

	return err
}
