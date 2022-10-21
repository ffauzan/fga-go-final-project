package sqldb

import (
	"final-project/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) domain.SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (r *SocialMediaRepository) SaveSocialMedia(socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	dbSocialMedia := SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
	}

	err := r.db.Create(&dbSocialMedia).Error
	if err != nil {
		return nil, err
	}

	socialMedia.ID = dbSocialMedia.ID
	socialMedia.CreatedAt = dbSocialMedia.CreatedAt
	socialMedia.UpdatedAt = dbSocialMedia.UpdatedAt

	return socialMedia, nil
}

func (r *SocialMediaRepository) UpdateSocialMedia(socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	err := r.db.Model(&SocialMedia{}).Where("id = ?", socialMedia.ID).Updates(SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}).Error

	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (r *SocialMediaRepository) GetSocialMediaByID(socialMediaID uint) (*domain.SocialMedia, error) {
	var dbSocialMedia SocialMedia
	err := r.db.First(&dbSocialMedia, socialMediaID).Error
	if err != nil {
		return nil, err
	}

	socialMedia := domain.SocialMedia{
		ID:             dbSocialMedia.ID,
		Name:           dbSocialMedia.Name,
		SocialMediaUrl: dbSocialMedia.SocialMediaUrl,
		UserID:         dbSocialMedia.UserID,
		CreatedAt:      dbSocialMedia.CreatedAt,
		UpdatedAt:      dbSocialMedia.UpdatedAt,
	}

	return &socialMedia, nil
}

func (r *SocialMediaRepository) GetSocialMediasByUserID(userID uint) (*[]domain.SocialMedia, error) {
	var dbSocialMedias []SocialMedia
	err := r.db.Where("user_id = ?", userID).Find(&dbSocialMedias).Error
	if err != nil {
		return nil, err
	}

	socialMedias := make([]domain.SocialMedia, len(dbSocialMedias))
	for i, dbSocialMedia := range dbSocialMedias {
		socialMedias[i] = domain.SocialMedia{
			ID:             dbSocialMedia.ID,
			Name:           dbSocialMedia.Name,
			SocialMediaUrl: dbSocialMedia.SocialMediaUrl,
			UserID:         dbSocialMedia.UserID,
			CreatedAt:      dbSocialMedia.CreatedAt,
			UpdatedAt:      dbSocialMedia.UpdatedAt,
		}
	}

	return &socialMedias, nil
}

func (r *SocialMediaRepository) DeleteSocialMediaByID(socialMediaID uint) error {
	err := r.db.Delete(&SocialMedia{}, socialMediaID).Error
	if err != nil {
		return err
	}

	return nil
}
