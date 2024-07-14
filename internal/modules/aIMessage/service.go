package aImessage

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type AIMessageService interface {
	SendMessage(senderId uint, conversationID uint, message string) (*entity.AIMessage, error)
	GetAllMessages(conversationID uint) (*[]entity.AIMessage, error)
	UpdateMessage(id uint, message string) (*entity.AIMessage, error)
	DeleteMessage(id uint) error
}

type aIMessageService struct {
	aIMessageRepository AIMessageRepository
}

func NewAIMessageService(aIMessageRepository AIMessageRepository) AIMessageService {
	return &aIMessageService{
		aIMessageRepository: aIMessageRepository,
	}
}

func (s *aIMessageService) SendMessage(senderId uint, conversationID uint, message string) (*entity.AIMessage, error) {

	aIMessage, err := s.aIMessageRepository.Create(senderId, conversationID, message, message)

	if err != nil {
		return nil, err
	}

	return aIMessage, nil
}

func (s *aIMessageService) GetAllMessages(conversationID uint) (*[]entity.AIMessage, error) {

	aIMessages, err := s.aIMessageRepository.FindAll(conversationID)

	if err != nil {
		return nil, err
	}

	return aIMessages, nil
}

func (s *aIMessageService) UpdateMessage(id uint, message string) (*entity.AIMessage, error) {

	aIMessage, err := s.aIMessageRepository.Update(id, message)

	if err != nil {
		return nil, err
	}

	return aIMessage, nil
}

func (s *aIMessageService) DeleteMessage(id uint) error {

	err := s.aIMessageRepository.DeleteOne(id)

	return err
}