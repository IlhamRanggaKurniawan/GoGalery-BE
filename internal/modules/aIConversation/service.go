package aIconversation

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type AIConversationService interface {
	CreateConversation(userId uint64) (*entity.AIConversation, error)
	GetConversation(userId uint64) (*entity.AIConversation, error)
	DeleteConversation(id uint64) error
}

type aIConversationService struct {
	aIConversationRepository AIConversationRepository
}

func NewAIConversationService(aIConversationRepository AIConversationRepository) AIConversationService {
	return &aIConversationService{
		aIConversationRepository: aIConversationRepository,
	}
}

func (s *aIConversationService) CreateConversation(userId uint64) (*entity.AIConversation, error) {

	directMessage, err := s.aIConversationRepository.Create(userId)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *aIConversationService) GetConversation(userId uint64) (*entity.AIConversation, error) {

	directMessage, err := s.aIConversationRepository.FindOne(userId)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *aIConversationService) DeleteConversation(id uint64) error{

	err := s.aIConversationRepository.DeleteOne(id)

	return err
}
