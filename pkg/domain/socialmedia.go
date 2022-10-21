package domain

import "time"

type SocialMedia struct {
	ID             uint
	Name           string
	SocialMediaUrl string
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SocialMediaService interface {
	AddSocialMedia(userID uint, name string, socialMediaUrl string) (*SocialMedia, error)
	GetSocialMediaByID(socialMediaID uint) (*SocialMedia, error)
	GetSocialMediasByUserID(userID uint) (*[]SocialMedia, error)
	UpdateSocialMedia(socialMediaID uint, name string, socialMediaUrl string) (*SocialMedia, error)
	DeleteSocialMedia(socialMediaID uint) error
}

type SocialMediaRepository interface {
	SaveSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error)
	GetSocialMediaByID(socialMediaID uint) (*SocialMedia, error)
	GetSocialMediasByUserID(userID uint) (*[]SocialMedia, error)
	UpdateSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error)
	DeleteSocialMediaByID(socialMediaID uint) error
}
