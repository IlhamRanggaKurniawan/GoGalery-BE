package notification

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type NotificationService interface {
	CreateNotification(receiverId uint64, triggerId uint64, content string) (*entity.Notification, error)
	GetAllNotifications(receiverId uint64) (*[]entity.Notification, error)
	UpdateNotifications(receiverId uint64) (*[]entity.Notification, error)
	DeleteNotifications(receiverId uint64) error
}

type notificationService struct {
	notificationRepository NotificationRepository
}

func NewNotificationService(notificationRepository NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
	}
}

func (s *notificationService) CreateNotification(receiverId uint64, triggerId uint64, content string) (*entity.Notification, error){

	notification, err := s.notificationRepository.Create(receiverId, triggerId, content)

	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (s *notificationService) GetAllNotifications(receiverId uint64) (*[]entity.Notification, error) {

	notifications, err := s.notificationRepository.FindAll(receiverId)

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (s *notificationService) UpdateNotifications(receiverId uint64) (*[]entity.Notification, error) {

	message, err := s.notificationRepository.Update(receiverId)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func (s *notificationService) DeleteNotifications(receiverId uint64) error {

	err := s.notificationRepository.DeleteAll(receiverId)

	return err
}