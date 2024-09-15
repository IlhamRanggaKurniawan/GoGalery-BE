package follow

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type FollowService interface {
	followUser(followerId uint64, followingId uint64) (*entity.Follow, error)
	GetAllFollower(userId uint64) (*[]entity.Follow, error)
	GetAllFollowing(userId uint64) (*[]entity.Follow, error)
	CountFollowing(userId uint64) (int64, error)
	CountFollower(userId uint64) (int64, error)
	CheckFollowing(followerId uint64, followingId uint64) (*entity.Follow, error)
	UnfollowUser(id uint64) error
}

type followService struct {
	followRepository FollowRepository
}

func NewFollowService(followRepository FollowRepository) FollowService {
	return &followService{
		followRepository: followRepository,
	}
}

func (s *followService) followUser(followerId uint64, followingId uint64) (*entity.Follow, error) {

	follow, err := s.followRepository.Create(followerId, followingId)

	if err != nil {
		return nil, err
	}

	return follow, nil
}

func (s *followService) GetAllFollower(userId uint64) (*[]entity.Follow, error) {

	follower, err := s.followRepository.FindAllFollower(userId)

	if err != nil {
		return nil, err
	}

	return follower, nil
}

func (s *followService) GetAllFollowing(userId uint64) (*[]entity.Follow, error) {

	following, err := s.followRepository.FindAllFollowing(userId)

	if err != nil {
		return nil, err
	}

	return following, nil
}

func (s *followService) CountFollowing(userId uint64) (int64, error) {

	following, err := s.followRepository.CountFollowing(userId)

	if err != nil {
		return 0, err
	}

	return following, nil
}

func (s *followService) CountFollower(userId uint64) (int64, error) {

	follower, err := s.followRepository.CountFollower(userId)

	if err != nil {
		return 0, err
	}

	return follower, nil
}

func (s *followService) CheckFollowing(followerId uint64, followingId uint64) (*entity.Follow, error) {

	follow, err := s.followRepository.FindOne(followerId, followingId)

	if err != nil {
		return nil, err
	}

	return follow, nil
}

func (s *followService) UnfollowUser(id uint64) error {

	err := s.followRepository.DeleteOne(id)

	return err
}
