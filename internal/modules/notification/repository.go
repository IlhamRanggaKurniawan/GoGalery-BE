package notification

import (

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(receiverId uint, triggerId uint, content string) (*entity.Notification, error)
	FindAll(receiverId uint) (*[]entity.Notification, error)
	Update(receiverId uint) (*[]entity.Notification, error)
	DeleteAll(receiverId uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(receiverId uint, triggerId uint, content string) (*entity.Notification, error) {
	notification := entity.Notification{
		Content:    content,
		ReceiverID: receiverId,
		TriggerID:  triggerId,
	}

	err := r.db.Create(&notification).Error

	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) FindAll(receiverId uint) (*[]entity.Notification, error) {
	var notifications []entity.Notification
	err := r.db.Where("receiver_id = ?", receiverId).Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	return &notifications, nil
}

func (r *notificationRepository) Update(receiverId uint) (*[]entity.Notification, error) {
	var notifications []entity.Notification

	err := r.db.Where("receiver_id = ? AND is_checked = ?", receiverId, false).Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	for i := range notifications {
		notifications[i].IsChecked = true
	}

	r.db.Save(&notifications)

	return &notifications, nil
}

func (r *notificationRepository) DeleteAll(receiverId uint) error {

	err := r.db.Where("receiver_id = ?", receiverId).Delete(&entity.Notification{}).Error

	if err != nil {
		return err
	}

	return nil
}
