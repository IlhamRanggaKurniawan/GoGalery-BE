package aImessage

import (

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	fetchapi "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/fetchAPI"
)

type AIMessageService interface {
	SendMessage(senderId uint64, conversationID uint64, prompt []fetchapi.Message) (*entity.AIMessage, error)
	GetAllMessages(conversationID uint64) (*[]entity.AIMessage, error)
	UpdateMessage(id uint64, message string) (*entity.AIMessage, error)
	DeleteMessage(id uint64) error
}

type aIMessageService struct {
	aIMessageRepository AIMessageRepository
}

func NewAIMessageService(aIMessageRepository AIMessageRepository) AIMessageService {
	return &aIMessageService{
		aIMessageRepository: aIMessageRepository,
	}
}

func (s *aIMessageService) SendMessage(senderId uint64, conversationID uint64, prompt []fetchapi.Message) (*entity.AIMessage, error) {

	response, err := fetchapi.FetchOpenAI(prompt)

	if err != nil {
		return nil, err
	}

	aIMessage, err := s.aIMessageRepository.Create(senderId, conversationID, prompt[len(prompt)-1].Content, response.Choices[0].Message.Content)
	if err != nil {
		return nil, err
	}
	return aIMessage, nil
}

func (s *aIMessageService) GetAllMessages(conversationID uint64) (*[]entity.AIMessage, error) {

	aIMessages, err := s.aIMessageRepository.FindAll(conversationID)

	if err != nil {
		return nil, err
	}

	return aIMessages, nil
}

func (s *aIMessageService) UpdateMessage(id uint64, message string) (*entity.AIMessage, error) {

	aIMessage, err := s.aIMessageRepository.Update(id, message)

	if err != nil {
		return nil, err
	}

	return aIMessage, nil
}

func (s *aIMessageService) DeleteMessage(id uint64) error {

	err := s.aIMessageRepository.DeleteOne(id)

	return err
}
