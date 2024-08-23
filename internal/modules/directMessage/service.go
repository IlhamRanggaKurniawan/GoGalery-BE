package directmessage

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type DirectMessageService interface {
	CreateDirectMessage(participants []uint64) (*entity.DirectMessage, error)
	GetAllDirectMessages(userId uint64) (*[]entity.DirectMessage, error)
	GetOneDirectMessage(id uint64) (*entity.DirectMessage, error)
	GetOneDirectMessageByParticipants(participants []uint64) (*entity.DirectMessage, error)
	DeleteDirectMessage(id uint64) error
}

type directMessageService struct {
	directMessageRepository DirectMessageRepository
}

func NewDirectMessageService(directMessageRepository DirectMessageRepository) DirectMessageService {
	return &directMessageService{
		directMessageRepository: directMessageRepository,
	}
}

func (s *directMessageService) CreateDirectMessage(participants []uint64) (*entity.DirectMessage, error) {

	directMessage, err := s.directMessageRepository.Create(participants)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *directMessageService) GetAllDirectMessages(userId uint64) (*[]entity.DirectMessage, error) {

	directMessages, err := s.directMessageRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return directMessages, nil
}

func (s *directMessageService) GetOneDirectMessageByParticipants(participants []uint64) (*entity.DirectMessage, error) {

	directMessage, err := s.directMessageRepository.FindOneByParticipants(participants)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *directMessageService) GetOneDirectMessage(id uint64) (*entity.DirectMessage, error) {

	directMessage, err := s.directMessageRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *directMessageService) DeleteDirectMessage(id uint64) error {

	err := s.directMessageRepository.DeleteOne(id)

	return err
}
