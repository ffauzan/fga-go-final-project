package photo

import (
	"errors"

	"final-project/pkg/domain"
)

type Repository interface {
	SavePhoto(*domain.Photo) (*domain.Photo, error)
	GetPhoto(uint) (*domain.Photo, error)
	GetPhotos() ([]domain.Photo, error)
	DeletePhoto(uint) error
	UpdatePhoto(*domain.Photo) (*domain.Photo, error)

	SaveComment(*domain.Comment) (*domain.Comment, error)
	GetComment(uint) (*domain.Comment, error)
	GetComments() ([]domain.Comment, error)
	DeleteComment(uint) error
	UpdateComment(*domain.Comment) (*domain.Comment, error)
}

type Service interface {
	SavePhoto(*domain.Photo) (*domain.Photo, error)
	GetPhoto(uint) (*domain.Photo, error)
	GetPhotos() ([]domain.Photo, error)
	DeletePhoto(uint) error
	UpdatePhoto(*domain.Photo) (*domain.Photo, error)
	IsPhotoOwner(uint, uint) bool
	IsPhotoExist(uint) bool

	SaveComment(*domain.Comment) (*domain.Comment, error)
	GetComment(uint) (*domain.Comment, error)
	GetComments() ([]domain.Comment, error)
	DeleteComment(uint) error
	UpdateComment(*domain.Comment) (*domain.Comment, error)
	IsCommentOwner(uint, uint) bool
	IsCommentExist(uint) bool
}

var (
	ErrNotFound     = errors.New("photo: resource not found")
	ErrUnauthorized = errors.New("photo: unauthorized")
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) SavePhoto(photo *domain.Photo) (*domain.Photo, error) {
	return s.repo.SavePhoto(photo)
}

func (s *service) GetPhoto(id uint) (*domain.Photo, error) {
	return s.repo.GetPhoto(id)
}

func (s *service) GetPhotos() ([]domain.Photo, error) {
	return s.repo.GetPhotos()
}

func (s *service) DeletePhoto(id uint) error {
	return s.repo.DeletePhoto(id)
}

func (s *service) UpdatePhoto(photo *domain.Photo) (*domain.Photo, error) {
	if !s.IsPhotoExist(photo.ID) {
		return nil, ErrNotFound
	}

	if !s.IsPhotoOwner(photo.ID, photo.UserID) {
		return nil, ErrUnauthorized
	}

	return s.repo.UpdatePhoto(photo)
}

func (s *service) IsPhotoOwner(photoID, userID uint) bool {
	photo, err := s.repo.GetPhoto(photoID)
	if err != nil {
		return false
	}

	if photo.UserID != userID {
		return false
	}

	return true
}

func (s *service) IsPhotoExist(photoID uint) bool {
	_, err := s.repo.GetPhoto(photoID)
	return err == nil
}

func (s *service) SaveComment(comment *domain.Comment) (*domain.Comment, error) {
	return s.repo.SaveComment(comment)
}

func (s *service) GetComment(id uint) (*domain.Comment, error) {
	return s.repo.GetComment(id)
}

func (s *service) GetComments() ([]domain.Comment, error) {
	return s.repo.GetComments()
}

func (s *service) DeleteComment(id uint) error {
	return s.repo.DeleteComment(id)
}

func (s *service) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {
	if !s.IsCommentExist(comment.ID) {
		return nil, ErrNotFound
	}

	if !s.IsCommentOwner(comment.ID, comment.UserID) {
		return nil, ErrUnauthorized
	}

	return s.repo.UpdateComment(comment)
}

func (s *service) IsCommentOwner(commentID, userID uint) bool {
	comment, err := s.repo.GetComment(commentID)
	if err != nil {
		return false
	}

	if comment.UserID != userID {
		return false
	}

	return true
}

func (s *service) IsCommentExist(commentID uint) bool {
	_, err := s.repo.GetComment(commentID)
	return err == nil
}
