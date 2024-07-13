package user

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username string, email string, password string) (*entity.User, error)
	Login(username string, password string) (*entity.User, error)
	UpdateUser(username *string, bio *string, profileUrl *string, password *string) (*entity.User, error)
	FindAllUsers(username string) (*[]entity.User, error)
	FindOneUser(username string) (*entity.User, error)
	DeleteUser(id uint) error
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	hashedPasswordStr := string(hashedPassword)

	user, err := s.userRepository.Create(username, email, hashedPasswordStr)

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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(username *string, bio *string, profileUrl *string, password *string) (*entity.User, error) {

	user, err := s.userRepository.Update(username, bio, profileUrl, password)

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

func (s *userService) DeleteUser(id uint) error {

	err := s.userRepository.DeleteOne(id)

	return err
}
