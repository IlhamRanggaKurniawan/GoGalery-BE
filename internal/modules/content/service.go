package content

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type ContentService interface {
	UploadContent(uploaderId *uint, caption *string, url *string) (*entity.Content, error)
	UpdateContent(id *uint, caption *string) (*entity.Content, error)
	GetAllContents() (*[]entity.Content, error)
	GetOneContent(id *uint) (*entity.Content, error)
	DeleteContent(id *uint) error
}

type contentService struct {
	contentRepository ContentRepository
}

func NewContentService(contentRepository ContentRepository) ContentService {
	return &contentService{
		contentRepository: contentRepository,
	}
}

func (s *contentService) UploadContent(uploaderId *uint, caption *string, url *string) (*entity.Content, error) {

	content, err := s.contentRepository.Create(uploaderId, caption, url)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *contentService) UpdateContent(id *uint, caption *string) (*entity.Content, error) {

	content, err := s.contentRepository.Update(id, caption)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *contentService) GetAllContents() (*[]entity.Content, error) {

	contents, err := s.contentRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (s *contentService) GetOneContent(id *uint) (*entity.Content, error) {

	content, err := s.contentRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *contentService) DeleteContent(id *uint) error {

	err := s.contentRepository.DeleteOne(id)

	return err
}
