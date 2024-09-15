package follow

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type FollowRepository interface {
	Create(followerId uint64, followingId uint64) (*entity.Follow, error)
	FindAllFollower(userId uint64) (*[]entity.Follow, error)
	FindAllFollowing(userId uint64) (*[]entity.Follow, error)
	CountFollower(userId uint64) (int64, error)
	CountFollowing(userId uint64) (int64, error)
	FindOne(followerId uint64, followingId uint64) (*entity.Follow, error)
	DeleteOne(id uint64) error
}

type followRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{db: db}
}

func (r *followRepository) Create(followerId uint64, followingId uint64) (*entity.Follow, error) {
	follow := entity.Follow{
		FollowerID:  followerId,
		FollowingID: followingId,
	}

	err := r.db.Create(&follow).Error

	if err != nil {
		return nil, err
	}

	return &follow, nil
}

func (r *followRepository) FindAllFollower(userId uint64) (*[]entity.Follow, error) {
	var follower []entity.Follow

	err := r.db.Where("following_id = ?", userId).Find(&follower).Error

	if err != nil {
		return nil, err
	}

	return &follower, nil
}

func (r *followRepository) FindAllFollowing(userId uint64) (*[]entity.Follow, error) {
	var following []entity.Follow

	err := r.db.Where("follower_id = ?", userId).Find(&following).Error

	if err != nil {
		return nil, err
	}

	return &following, nil
}

func (r *followRepository) CountFollower(userId uint64) (int64, error) {
	var follower int64

	err := r.db.Model(&entity.Follow{}).Where("following_id = ?", userId).Count(&follower).Error

	return follower, err
}

func (r *followRepository) CountFollowing(userId uint64) (int64, error) {
	var following int64

	err := r.db.Model(&entity.Follow{}).Where("follower_id = ?", userId).Count(&following).Error

	return following, err
}

func (r *followRepository) FindOne(followerId uint64, followingId uint64) (*entity.Follow, error) {
	var follow entity.Follow

	err := r.db.Where("follower_id = ? AND following_id = ?", followerId, followingId).Take(&follow).Error

	if err != nil {
		return nil, err
	}

	return &follow, nil
}

func (r *followRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.Follow{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
