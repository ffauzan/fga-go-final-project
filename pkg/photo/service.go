package photo

import (
	"final-project/pkg/domain"
)

type service struct {
	repo domain.PhotoRepository
}

func NewService(repo domain.PhotoRepository) domain.PhotoService {
	return &service{
		repo: repo,
	}
}

func (s *service) SavePhoto(userID uint, photo *domain.AddPhotoRequest) (*domain.Photo, error) {
	photoToSave := &domain.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   userID,
	}
	return s.repo.SavePhoto(photoToSave)
}

func (s *service) GetPhotoByID(photoID uint) (*domain.Photo, error) {
	return s.repo.GetPhotoById(photoID)
}

func (s *service) GetPhotosByUserID(userID uint) (*[]domain.Photo, error) {
	return s.repo.GetPhotosByUserID(userID)
}

func (s *service) UpdatePhoto(photoID uint, newPhoto *domain.AddPhotoRequest) (*domain.Photo, error) {
	photo, err := s.GetPhotoByID(photoID)
	if err != nil {
		return nil, err
	}
	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoUrl = newPhoto.PhotoUrl
	return s.repo.UpdatePhoto(photo)
}

func (s *service) DeletePhoto(photoID uint) error {
	return s.repo.DeletePhotoById(photoID)
}
