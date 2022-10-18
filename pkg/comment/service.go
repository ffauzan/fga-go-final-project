package comment

import "errors"

type Repository interface {
	SaveComment(comment *Comment) (*Comment, error)
	GetCommentByID(commentID uint) (*Comment, error)
	GetComments() ([]Comment, error)
	DeleteComment(commentID uint) error
	UpdateComment(comment *Comment) (*Comment, error)
	GetUserByID(userID uint) (*User, error)
	GetPhotoByID(photoID uint) (*Photo, error)
}

type AuthService interface {
	IsUserCanCreate(userID uint, entity interface{}) bool
	IsUserCanAccess(userID uint, entity interface{}) bool
	IsUserCanUpdate(userID uint, entity interface{}) bool
	IsUserCanDelete(userID uint, entity interface{}) bool
}

type Service interface {
	AddComment(userID uint, photoID uint, message string) (*Comment, error)
	GetCommentsOfUser(userID uint) ([]Comment, error)
	UpdateComment(userID uint, commentID uint, message string) (*Comment, error)
	DeleteComment(userID uint, commentID uint) error
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

func (s *service) AddComment(userID uint, photoID uint, message string) (*Comment, error) {
	// Check if user exist
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Check if photo exist
	_, err = s.repo.GetPhotoByID(photoID)
	if err != nil {
		return nil, err
	}

	// Create comment
	comment := &Comment{
		UserID:  userID,
		PhotoID: photoID,
		Message: message,
	}

	// Save comment
	return s.repo.SaveComment(comment)
}

func (s *service) GetCommentsOfUser(userID uint) ([]Comment, error) {
	// Check if user exist
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Get comments
	return s.repo.GetComments()
}

func (s *service) UpdateComment(userID uint, commentID uint, message string) (*Comment, error) {
	// Check if user exist
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Get comment
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}

	// Check if user can update comment
	if !s.authService.IsUserCanUpdate(userID, comment) {
		return nil, errors.New("user can't update comment")
	}

	// Update comment
	comment.Message = message

	// Save comment
	return s.repo.UpdateComment(comment)
}

func (s *service) DeleteComment(userID uint, commentID uint) error {
	// Check if user exist
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Get comment
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return err
	}

	// Check if user can delete comment
	if !s.authService.IsUserCanDelete(userID, comment) {
		return errors.New("user can't delete comment")
	}

	// Delete comment
	return s.repo.DeleteComment(commentID)
}
