package directmessage

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type DirectMessageService interface {
	CreateDirectMessage(participants []entity.User) (*entity.DirectMessage, error)
	GetAllDirectMessages(userId uint) (*[]entity.DirectMessage, error)
	GetOneDirectMessage(id uint) (*entity.DirectMessage, error)
	DeleteDirectMessage(id uint) error
}

type directMessageService struct {
	directMessageRepository DirectMessageRepository
}

func NewDirectMessageService(directMessageRepository DirectMessageRepository) DirectMessageService {
	return &directMessageService{
		directMessageRepository: directMessageRepository,
	}
}

func (s *directMessageService) CreateDirectMessage(participants []entity.User) (*entity.DirectMessage, error) {

	directMessage, err := s.directMessageRepository.Create(participants)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *directMessageService) GetAllDirectMessages(userId uint) (*[]entity.DirectMessage, error) {

	directMessages, err := s.directMessageRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return directMessages, nil
}

func (s *directMessageService) GetOneDirectMessage(id uint) (*entity.DirectMessage, error) {

	directMessage, err := s.directMessageRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *directMessageService) DeleteDirectMessage(id uint) error {

	err := s.directMessageRepository.DeleteOne(id)

	return err
}
