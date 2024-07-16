package follow

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type FollowRepository interface {
	Create(followerId uint64, followingId uint64) (*entity.Follow, error)
	FindAll(userId uint64) (*[]entity.Follow, *[]entity.Follow, error)
	FindOne(followerId uint64, followingId uint64) (*entity.Follow, error)
	GetDB() *gorm.DB
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

func (r *followRepository) FindAll(userId uint64) (*[]entity.Follow, *[]entity.Follow, error) {
	var follower []entity.Follow
	var following []entity.Follow

	err := r.db.Where("following_id = ?", userId).Find(&follower).Error

	if err != nil {
		return nil, nil, err
	}

	err = r.db.Where("follower_id = ?", userId).Find(&following).Error

	if err != nil {
		return nil, nil, err
	}

	return &follower, &following, nil
}

func (r *followRepository) FindOne(followerId uint64, followingId uint64) (*entity.Follow, error) {
	var follow entity.Follow

	err := r.db.Where("follower_id = ? AND following_id = ?", followerId, followingId).Take(&follow).Error

	if err != nil {
		return nil, err
	}

	return &follow, nil
}

func (r *followRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *followRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.Follow{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
