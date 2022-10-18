package photo

import (
	"errors"
)

type Repository interface {
	SavePhoto(*Photo) (*Photo, error)
	GetPhoto(uint) (*Photo, error)
	GetPhotos() ([]Photo, error)
	DeletePhoto(uint) error
	UpdatePhoto(*Photo) (*Photo, error)

	SaveComment(*Comment) (*Comment, error)
	GetComment(uint) (*Comment, error)
	GetComments() ([]Comment, error)
	DeleteComment(uint) error
	UpdateComment(*Comment) (*Comment, error)
}

type Service interface {
	SavePhoto(*Photo) (*Photo, error)
	GetPhoto(uint) (*Photo, error)
	GetPhotos() ([]Photo, error)
	DeletePhoto(uint) error
	UpdatePhoto(*Photo) (*Photo, error)
	IsPhotoOwner(uint, uint) bool
	IsPhotoExist(uint) bool

	SaveComment(*Comment) (*Comment, error)
	GetComment(uint) (*Comment, error)
	GetComments() ([]Comment, error)
	DeleteComment(uint) error
	UpdateComment(*Comment) (*Comment, error)
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

func (s *service) SavePhoto(photo *Photo) (*Photo, error) {
	return s.repo.SavePhoto(photo)
}

func (s *service) GetPhoto(id uint) (*Photo, error) {
	return s.repo.GetPhoto(id)
}

func (s *service) GetPhotos() ([]Photo, error) {
	return s.repo.GetPhotos()
}

func (s *service) DeletePhoto(id uint) error {
	return s.repo.DeletePhoto(id)
}

func (s *service) UpdatePhoto(photo *Photo) (*Photo, error) {
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

func (s *service) SaveComment(comment *Comment) (*Comment, error) {
	return s.repo.SaveComment(comment)
}

func (s *service) GetComment(id uint) (*Comment, error) {
	return s.repo.GetComment(id)
}

func (s *service) GetComments() ([]Comment, error) {
	return s.repo.GetComments()
}

func (s *service) DeleteComment(id uint) error {
	return s.repo.DeleteComment(id)
}

func (s *service) UpdateComment(comment *Comment) (*Comment, error) {
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
