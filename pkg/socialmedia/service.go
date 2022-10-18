package socialmedia

import "errors"

type Repository interface {
	SaveSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error)
	GetSocialMediaByID(socialMediaID uint) (*SocialMedia, error)
	GetSocialMediasOfUser(userId uint) (*[]SocialMedia, error)
	DeleteSocialMedia(socialMediaID uint) error
	UpdateSocialMedia(socialMedia *SocialMedia) (*SocialMedia, error)
}

type AuthService interface {
	IsUserCanCreate(userID uint, entity interface{}) bool
	IsUserCanAccess(userID uint, entity interface{}) bool
	IsUserCanUpdate(userID uint, entity interface{}) bool
	IsUserCanDelete(userID uint, entity interface{}) bool
}

type Service interface {
	AddSocialMedia(userID uint, req *AddSocialMediaRequest) (*SocialMedia, error)
	GetSocialMedia(socialMediaID uint) (*SocialMedia, error)
	GetSocialMedias(userID uint) (*[]SocialMedia, error)
	UpdateSocialMedia(userID uint, socialMediaID uint, req AddSocialMediaRequest) (*SocialMedia, error)
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

func (s *service) AddSocialMedia(userID uint, req *AddSocialMediaRequest) (*SocialMedia, error) {
	// create social media
	socialMedia := &SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         userID,
	}

	return s.repo.SaveSocialMedia(socialMedia)
}

func (s *service) GetSocialMedia(socialMediaID uint) (*SocialMedia, error) {
	return s.repo.GetSocialMediaByID(socialMediaID)
}

func (s *service) GetSocialMedias(userID uint) (*[]SocialMedia, error) {
	return s.repo.GetSocialMediasOfUser(userID)
}

func (s *service) UpdateSocialMedia(userID uint, socialMediaID uint, req AddSocialMediaRequest) (*SocialMedia, error) {
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
