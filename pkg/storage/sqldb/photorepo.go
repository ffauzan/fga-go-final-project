package sqldb

import (
	"final-project/pkg/domain"
	"log"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment `gorm:"foreignKey:PhotoID"`
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	PhotoID   uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	log.Println("PhotoRepository created")
	return &PhotoRepository{
		db: db,
	}
}

func (r *PhotoRepository) SavePhoto(photo *domain.Photo) (*domain.Photo, error) {
	dbPhoto := Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   photo.UserID,
	}

	err := r.db.Create(&dbPhoto).Error
	if err != nil {
		return nil, err
	}

	photo.ID = dbPhoto.ID
	photo.CreatedAt = dbPhoto.CreatedAt
	photo.UpdatedAt = dbPhoto.UpdatedAt

	return photo, nil
}

func (r *PhotoRepository) GetPhotoByID(photoID uint) (*domain.Photo, error) {
	var dbPhoto Photo
	err := r.db.First(&dbPhoto, photoID).Error
	if err != nil {
		return nil, err
	}

	photo := domain.Photo{
		ID:        dbPhoto.ID,
		Title:     dbPhoto.Title,
		Caption:   dbPhoto.Caption,
		PhotoUrl:  dbPhoto.PhotoUrl,
		UserID:    dbPhoto.UserID,
		CreatedAt: dbPhoto.CreatedAt,
		UpdatedAt: dbPhoto.UpdatedAt,
	}

	return &photo, nil
}

func (r *PhotoRepository) GetPhotosByUserID(userId uint) (*[]domain.Photo, error) {
	var dbPhotos []Photo
	err := r.db.Where("user_id = ?", userId).Find(&dbPhotos).Error
	if err != nil {
		return nil, err
	}

	photos := make([]domain.Photo, len(dbPhotos))
	for i, dbPhoto := range dbPhotos {
		photos[i] = domain.Photo{
			ID:        dbPhoto.ID,
			Title:     dbPhoto.Title,
			Caption:   dbPhoto.Caption,
			PhotoUrl:  dbPhoto.PhotoUrl,
			UserID:    dbPhoto.UserID,
			CreatedAt: dbPhoto.CreatedAt,
			UpdatedAt: dbPhoto.UpdatedAt,
		}
	}

	return &photos, nil
}

func (r *PhotoRepository) UpdatePhoto(photo *domain.Photo) (*domain.Photo, error) {
	err := r.db.Model(Photo{}).Where("id = ?", photo.ID).Updates(Photo{
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UpdatedAt: time.Now(),
	}).Error

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *PhotoRepository) DeletePhotoByID(photoID uint) error {
	err := r.db.Delete(&Photo{}, photoID).Error
	if err != nil {
		return err
	}

	return nil
}
