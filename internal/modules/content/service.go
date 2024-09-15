package content

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type ContentService interface {
	UploadContent(uploaderId uint64, caption string, url string, contentType entity.ContentType) (*entity.Content, error)
	UpdateContent(id uint64, caption string) (*entity.Content, error)
	GetAllContents() (*[]entity.Content, error)
	GetAllContentsByFollowing(userId uint64) (*[]entity.Content, error)
	GetOneContent(id uint64) (*entity.Content, error)
	DeleteContent(id uint64) error
}

type contentService struct {
	contentRepository ContentRepository
}

func NewContentService(contentRepository ContentRepository) ContentService {
	return &contentService{
		contentRepository: contentRepository,
	}
}

func (s *contentService) UploadContent(uploaderId uint64, caption string, url string, contentType entity.ContentType) (*entity.Content, error) {

	content, err := s.contentRepository.Create(uploaderId, caption, url, contentType)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *contentService) UpdateContent(id uint64, caption string) (*entity.Content, error) {

	content, err := s.contentRepository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	content.Caption = caption

	content, err = s.contentRepository.Update(content)

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

func (s *contentService) GetAllContentsByFollowing(userId uint64) (*[]entity.Content, error) {

	contents, err := s.contentRepository.FindAllByFollowing(userId)

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (s *contentService) GetOneContent(id uint64) (*entity.Content, error) {

	content, err := s.contentRepository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *contentService) DeleteContent(id uint64) error {

	err := s.contentRepository.DeleteOne(id)

	return err
}
