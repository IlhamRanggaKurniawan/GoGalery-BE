package user

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(username string, email string, password string) (*entity.User, error)
	FindAllMutualUsers(userId uint64) (*[]entity.User, error)
	FindAllByUsername(username string) (*[]entity.User, error)
	FindOneByUsername(username string) (*entity.User, error)
	FindOneById(id uint64) (*entity.User, error)
	FindOneByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	DeleteOne(id uint64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(username string, email string, password string) (*entity.User, error) {
	user := entity.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := r.db.Create(&user).Error

	if err != nil {
		return nil, fmt.Errorf(`"Username" or "Email" already used`)
	}

	return &user, nil
}

func (r *userRepository) FindAllMutualUsers(userId uint64) (*[]entity.User, error) {
	var users []entity.User

	err := r.db.Table("users").
		Select("users.*").
		Joins("JOIN follows f1 ON f1.follower_id = users.id").
		Joins("JOIN follows f2 ON f2.following_id = users.id AND f2.follower_id = ?", userId).
		Where("f1.following_id = ?", userId).
		Scan(&users).Error

	return &users, err
}

func (r *userRepository) FindAllByUsername(username string) (*[]entity.User, error) {
	var users []entity.User

	err := r.db.Where("username LIKE ?", "%"+username+"%").Find(&users).Error

	return &users, err
}

func (r *userRepository) FindOneByUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("username = ?", username).Preload("Contents").Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneById(id uint64) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *entity.User) (*entity.User, error) {

	r.db.Save(&user)

	return user, nil
}

func (r *userRepository) DeleteOne(id uint64) error {
	err := r.db.Delete(&entity.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
