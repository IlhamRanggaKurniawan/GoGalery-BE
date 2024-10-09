package user

import (
	"fmt"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type UserService interface {
	Register(username string, email string, password string) (*entity.User, error)
	Login(username string, password string) (*entity.User, error)
	UpdateUser(id uint64, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error)
	UpdateUserByEmail(email string, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error)
	FindAllUsersByUsername(username string) (*[]entity.User, error)
	FindAllMutualUsers(userId uint64) (*[]entity.User, error)
	FindOneUserByUsername(username string) (*entity.User, error)
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

	user, err := s.userRepository.FindOneByUsername(username)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, password)

	if err != nil {
		return nil, fmt.Errorf("the provided credentials do not match our records")
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint64, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error) {

	user, err := s.userRepository.FindOneById(id)

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

	user, err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUserByEmail(email string, bio *string, profileUrl *string, password *string, token *string) (*entity.User, error) {

	user, err := s.userRepository.FindOneByEmail(email)

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

	user, err = s.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) FindAllUsersByUsername(username string) (*[]entity.User, error) {

	users, err := s.userRepository.FindAllByUsername(username)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindAllMutualUsers(userId uint64) (*[]entity.User, error) {

	users, err := s.userRepository.FindAllMutualUsers(userId)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindOneUserByUsername(username string) (*entity.User, error) {

	user, err := s.userRepository.FindOneByUsername(username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint64) error {

	err := s.userRepository.DeleteOne(id)

	return err
}
