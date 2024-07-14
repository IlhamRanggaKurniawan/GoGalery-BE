package groupchat

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type GroupChatService interface {
	CreateGroupChat(name string, members []entity.User) (*entity.GroupChat, error)
	GetAllGroupChats(userId uint) (*[]entity.GroupChat, error)
	GetOneGroupChat(id uint) (*entity.GroupChat, error)
	UpdateGroupChat(id uint, pictureUrl string) (*entity.GroupChat, error)
	DeleteGroupChat(id uint) error
}

type groupChatService struct {
	groupChatRepository GroupChatRepository
}

func NewGroupChatService(groupChatRepository GroupChatRepository) GroupChatService {
	return &groupChatService{
		groupChatRepository: groupChatRepository,
	}
}

func (s *groupChatService) CreateGroupChat(name string,members []entity.User) (*entity.GroupChat, error) {

	groupChat, err := s.groupChatRepository.Create(name,members)

	if err != nil {
		return nil, err
	}

	return groupChat, nil
}

func (s *groupChatService) GetAllGroupChats(userId uint) (*[]entity.GroupChat, error) {

	directMessages, err := s.groupChatRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return directMessages, nil
}

func (s *groupChatService) GetOneGroupChat(id uint) (*entity.GroupChat, error) {

	directMessage, err := s.groupChatRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *groupChatService) UpdateGroupChat(id uint, pictureUrl string) (*entity.GroupChat, error) {

	directMessage, err := s.groupChatRepository.Update(id, pictureUrl)

	if err != nil {
		return nil, err
	}

	return directMessage, nil
}

func (s *groupChatService) DeleteGroupChat(id uint) error {

	err := s.groupChatRepository.DeleteOne(id)

	return err
}
