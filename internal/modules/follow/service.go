package follow

import "github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"

type FollowService interface {
	followUser(followerId uint64, followingId uint64) (*entity.Follow, error)
	GetAllFollows(userId uint64) (*[]entity.Follow, *[]entity.Follow, error)
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

func (s *followService) GetAllFollows(userId uint64) (*[]entity.Follow, *[]entity.Follow, error) {

	follower, following, err := s.followRepository.FindAll(userId)

	if err != nil {
		return nil, nil, err
	}

	return follower, following, nil
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
