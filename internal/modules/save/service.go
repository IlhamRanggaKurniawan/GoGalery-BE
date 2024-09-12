package save

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type SaveContentService interface {
	SaveContent(userId uint64, contentId uint64) (*entity.SaveContent, error)
	GetAllSaves(userId uint64) (*[]entity.SaveContent, error)
	GetOneSave(userId uint64, contentId uint64) (*entity.SaveContent, error)
	UnsaveContent(id uint64) error
}

type saveContentService struct {
	saveContentRepository SaveContentRepository
}

func NewSaveService(saveContentRepository SaveContentRepository) SaveContentService {
	return &saveContentService{
		saveContentRepository: saveContentRepository,
	}
}

func (s *saveContentService) SaveContent(userId uint64, contentId uint64) (*entity.SaveContent, error) {

	save, err := s.saveContentRepository.Create(userId, contentId)

	if err != nil {
		return nil, err
	}

	return save, nil
}

func (s *saveContentService) GetAllSaves(userId uint64) (*[]entity.SaveContent, error) {

	saves, err := s.saveContentRepository.FindAllById(userId)

	if err != nil {
		return nil, err
	}

	return saves, nil
}

func (s *saveContentService) GetOneSave(userId uint64, contentId uint64) (*entity.SaveContent, error) {

	save, err := s.saveContentRepository.FindOneById(userId, contentId)

	if err != nil {
		return nil, err
	}

	return save, nil
}

func (s *saveContentService) UnsaveContent(id uint64) error {

	err := s.saveContentRepository.DeleteOne(id)

	return err
}
