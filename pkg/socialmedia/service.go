package socialmedia

import (
	"errors"
	"final-project/pkg/domain"
)

type AddSocialMediaRequest struct {
	Name           string
	SocialMediaUrl string
}

type Repository interface {
	SaveSocialMedia(socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	GetSocialMediaByID(socialMediaID uint) (*domain.SocialMedia, error)
	GetSocialMediasOfUser(userId uint) (*[]domain.SocialMedia, error)
	DeleteSocialMedia(socialMediaID uint) error
	UpdateSocialMedia(socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
}

type AuthService interface {
	IsUserCanCreate(userID uint, entity interface{}) bool
	IsUserCanAccess(userID uint, entity interface{}) bool
	IsUserCanUpdate(userID uint, entity interface{}) bool
	IsUserCanDelete(userID uint, entity interface{}) bool
}

type Service interface {
	AddSocialMedia(userID uint, req *AddSocialMediaRequest) (*domain.SocialMedia, error)
	GetSocialMedia(socialMediaID uint) (*domain.SocialMedia, error)
	GetSocialMedias(userID uint) (*[]domain.SocialMedia, error)
	UpdateSocialMedia(userID uint, socialMediaID uint, req AddSocialMediaRequest) (*domain.SocialMedia, error)
	DeleteSocialMedia(userID uint, socialMediaID uint) error

	// IsSocialMediaExist(uint) bool
}

type service struct {
	repo        Repository
	authService AuthService
}

func NewService(repo Repository, authService AuthService) Service {
	return &service{
		repo:        repo,
		authService: authService,
	}
}

func (s *service) AddSocialMedia(userID uint, req *AddSocialMediaRequest) (*domain.SocialMedia, error) {
	// create social media
	socialMedia := &domain.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         userID,
	}

	return s.repo.SaveSocialMedia(socialMedia)
}

func (s *service) GetSocialMedia(socialMediaID uint) (*domain.SocialMedia, error) {
	return s.repo.GetSocialMediaByID(socialMediaID)
}

func (s *service) GetSocialMedias(userID uint) (*[]domain.SocialMedia, error) {
	return s.repo.GetSocialMediasOfUser(userID)
}

func (s *service) UpdateSocialMedia(userID uint, socialMediaID uint, req AddSocialMediaRequest) (*domain.SocialMedia, error) {
	// get social media
	socialMedia, err := s.repo.GetSocialMediaByID(socialMediaID)
	if err != nil {
		return nil, err
	}

	// check if user can update social media
	if userID != socialMedia.UserID {
		return nil, errors.New("user can't update social media")
	}

	// update social media
	socialMedia.Name = req.Name
	socialMedia.SocialMediaUrl = req.SocialMediaUrl

	return s.repo.UpdateSocialMedia(socialMedia)
}

func (s *service) DeleteSocialMedia(userID uint, socialMediaID uint) error {
	// get social media
	socialMedia, err := s.repo.GetSocialMediaByID(socialMediaID)
	if err != nil {
		return err
	}

	// check if user can delete social media
	if userID != socialMedia.UserID {
		return errors.New("user can't delete social media")
	}

	return s.repo.DeleteSocialMedia(socialMediaID)
}
