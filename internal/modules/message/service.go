package message

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type MessageService interface {
	SendMessage(senderId uint, directMessageId uint, groupChatId uint, text string) (*entity.Message, error)
	GetAllMessages(directMessageId uint, groupChatId uint) (*[]entity.Message, error)
	UpdateMessage(id uint, text string) (*entity.Message, error)
	DeleteMessage(id uint) error
}

type messageService struct {
	messageRepository MessageRepository
}

func NewMessageService(messageRepository MessageRepository) MessageService {
	return &messageService{
		messageRepository: messageRepository,
	}
}

func (s *messageService) SendMessage(senderId uint, directMessageId uint, groupChatId uint, text string) (*entity.Message, error) {

	message, err := s.messageRepository.Create(senderId, directMessageId, groupChatId, text)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *messageService) GetAllMessages(directMessageId uint, groupChatId uint) (*[]entity.Message, error) {

	messages, err := s.messageRepository.FindAll(directMessageId, groupChatId)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *messageService) UpdateMessage(id uint, text string) (*entity.Message, error) {

	message, err := s.messageRepository.Update(id, text)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *messageService) DeleteMessage(id uint) error {

	err := s.messageRepository.DeleteOne(id)

	return err
}
