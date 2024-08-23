package user

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(username string, email string, password string) (*entity.User, error)
	FindAll(username string) (*[]entity.User, error)
	FindOne(username string) (*entity.User, error)
	FindOneByToken(token string) (*entity.User, error)
	Update(username string, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error)
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
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAll(username string) (*[]entity.User, error) {
	var users []entity.User

	err := r.db.Where("username LIKE ?", "%"+username+"%").Find(&users).Error

	return &users, err
}

func (r *userRepository) FindOne(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("username = ?", username).Preload("Contents").Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneByToken(token string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("token = ?", token).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(username string, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error) {

	user, err := r.FindOne(username)

	if err != nil {
		return nil, err
	}

	if bio != nil {
		user.Bio = bio
	}

	if profileUrl != nil {
		user.ProfileUrl = profileUrl
	}

	if token != nil {
		user.Token = token
	}
	
	if password != nil {
		hashedPassword, _ := utils.HashPassword(*password)

		user.Password = *hashedPassword
	}

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
