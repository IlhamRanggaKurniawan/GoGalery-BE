package save

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type SaveContentService interface {
	SaveContent(userId uint, contentId uint) (*entity.SaveContent, error)
	GetAllSaves(contentId uint) (*[]entity.SaveContent, error)
	GetOneSave(userId uint, contentId uint) (*entity.SaveContent, error)
	UnsaveContent(id uint) error
}

type saveContentService struct {
	saveContentRepository SaveContentRepository
}

func NewSaveService(saveContentRepository SaveContentRepository) SaveContentService {
	return &saveContentService{
		saveContentRepository: saveContentRepository,
	}
}

func (s *saveContentService) SaveContent(userId uint, contentId uint) (*entity.SaveContent, error) {

	save, err := s.saveContentRepository.Create(userId, contentId)

	if err != nil {
		return nil, err
	}

	return save, nil
}

func (s *saveContentService) GetAllSaves(contentId uint) (*[]entity.SaveContent, error) {

	saves, err := s.saveContentRepository.FindAll(contentId)

	if err != nil {
		return nil, err
	}

	return saves, nil
}

func (s *saveContentService) GetOneSave(userId uint, contentId uint) (*entity.SaveContent, error) {

	save, err := s.saveContentRepository.FindOne(userId, contentId)

	if err != nil {
		return nil, err
	}

	return save, nil
}

func (s *saveContentService) UnsaveContent(id uint) error {

	err := s.saveContentRepository.DeleteOne(id)

	return err
}
