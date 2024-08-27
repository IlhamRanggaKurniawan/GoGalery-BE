package user

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type UserService interface {
	Register(username string, email string, password string) (*entity.User, error)
	Login(username string, password string) (*entity.User, error)
	UpdateUser(id uint64, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error)
	FindAllUsers(username string) (*[]entity.User, error)
	FindOneUser(username string) (*entity.User, error)
	DeleteUser(id uint64) error
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Register(username string, email string, password string) (*entity.User, error) {

	hashedPassword, _ := utils.HashPassword(password)

	user, err := s.userRepository.Create(username, email, *hashedPassword)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(username string, password string) (*entity.User, error) {

	user, err := s.userRepository.FindOne(username)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint64, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error) {

	user, err := s.userRepository.Update(id, bio, profileUrl, password, token)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) FindAllUsers(username string) (*[]entity.User, error) {

	users, err := s.userRepository.FindAll(username)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindOneUser(username string) (*entity.User, error) {

	user, err := s.userRepository.FindOne(username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint64) error {

	err := s.userRepository.DeleteOne(id)

	return err
}
