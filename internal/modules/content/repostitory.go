package content

import (
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"gorm.io/gorm"
)

type ContentRepository interface {
	Create(uploaderId uint64, caption string, url string, contentType entity.ContentType) (*entity.Content, error)
	FindAll(limit int, nextCursor *uint64) (*[]entity.Content, *uint64, error)
	FindAllByFollowing(userId uint64) (*[]entity.Content, error)
	FindOneById(id uint64) (*entity.Content, error)
	Update(content *entity.Content) (*entity.Content, error)
	DeleteOne(id uint64) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{db: db}
}

func (r *contentRepository) Create(uploaderId uint64, caption string, url string, contentType entity.ContentType) (*entity.Content, error) {
	content := entity.Content{
		UploaderId: uploaderId,
		Caption:    caption,
		URL:        url,
		Type:       contentType,
	}

	err := r.db.Create(&content).Error

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r *contentRepository) FindAll(limit int, nextCursor *uint64) (*[]entity.Content, *uint64, error) {
	var contents []entity.Content

	query := r.db.Preload("Uploader").Order("id ASC").Limit(limit)

	if nextCursor != nil {
		query.Where("id > ?", *nextCursor)
	}

	err := query.Find(&contents).Error

	if err != nil {
		return nil, nil, err
	}

	if len(contents) == limit {
		nextCursor = &contents[limit-1].Id
	} else {
		nextCursor = nil
	}

	return &contents, nextCursor, nil
}

func (r *contentRepository) FindOneById(id uint64) (*entity.Content, error) {
	var content entity.Content

	err := r.db.Preload("Comments").Preload("Uploader").Where("id = ?", id).Take(&content).Error

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r *contentRepository) FindAllByFollowing(userId uint64) (*[]entity.Content, error) {
	var contents []entity.Content

	err := r.db.Joins("JOIN follows ON follows.following_id = contents.uploader_id").
		Where("follows.follower_id = ?", userId).
		Preload("Uploader").Find(&contents).Error

	if err != nil {
		return nil, err
	}

	return &contents, nil
}

func (r *contentRepository) Update(content *entity.Content) (*entity.Content, error) {

	r.db.Save(&content)

	return content, nil
}

func (r *contentRepository) DeleteOne(id uint64) error {

	err := r.db.Delete(&entity.Content{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
