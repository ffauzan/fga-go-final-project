package socialmedia

import (
	"final-project/pkg/domain"
)

type service struct {
	repo domain.SocialMediaRepository
}

func NewService(repo domain.SocialMediaRepository) domain.SocialMediaService {
	return &service{
		repo: repo,
	}
}

func (s *service) AddSocialMedia(userID uint, name string, socialMediaUrl string) (*domain.SocialMedia, error) {
	socialMedia := &domain.SocialMedia{
		UserID:         userID,
		Name:           name,
		SocialMediaUrl: socialMediaUrl,
	}

	return s.repo.SaveSocialMedia(socialMedia)
}

func (s *service) GetSocialMediaByID(socialMediaID uint) (*domain.SocialMedia, error) {
	return s.repo.GetSocialMediaByID(socialMediaID)
}

func (s *service) GetSocialMediasByUserID(userID uint) (*[]domain.SocialMedia, error) {
	return s.repo.GetSocialMediasByUserID(userID)
}

func (s *service) UpdateSocialMedia(socialMediaID uint, name string, socialMediaUrl string) (*domain.SocialMedia, error) {
	socialMedia, err := s.repo.GetSocialMediaByID(socialMediaID)
	if err != nil {
		return nil, err
	}

	socialMedia.Name = name
	socialMedia.SocialMediaUrl = socialMediaUrl

	return s.repo.UpdateSocialMedia(socialMedia)
}

func (s *service) DeleteSocialMedia(socialMediaID uint) error {
	return s.repo.DeleteSocialMediaByID(socialMediaID)
}
