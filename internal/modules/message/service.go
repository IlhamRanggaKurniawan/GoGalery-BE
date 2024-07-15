package message

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type MessageService interface {
	SendMessage(senderId uint64, directMessageId uint64, groupChatId uint64, text string) (*entity.Message, error)
	GetAllMessages(directMessageId uint64, groupChatId uint64) (*[]entity.Message, error)
	UpdateMessage(id uint64, text string) (*entity.Message, error)
	DeleteMessage(id uint64) error
}

type messageService struct {
	messageRepository MessageRepository
}

func NewMessageService(messageRepository MessageRepository) MessageService {
	return &messageService{
		messageRepository: messageRepository,
	}
}

func (s *messageService) SendMessage(senderId uint64, directMessageId uint64, groupChatId uint64, text string) (*entity.Message, error) {

	message, err := s.messageRepository.Create(senderId, directMessageId, groupChatId, text)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *messageService) GetAllMessages(directMessageId uint64, groupChatId uint64) (*[]entity.Message, error) {

	messages, err := s.messageRepository.FindAll(directMessageId, groupChatId)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *messageService) UpdateMessage(id uint64, text string) (*entity.Message, error) {

	message, err := s.messageRepository.Update(id, text)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *messageService) DeleteMessage(id uint64) error {

	err := s.messageRepository.DeleteOne(id)

	return err
}
