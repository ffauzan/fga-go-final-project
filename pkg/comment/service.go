package comment

import (
	"final-project/pkg/domain"
)

type service struct {
	repo domain.CommentRepository
}

func NewService(repo domain.CommentRepository) domain.CommentService {
	return &service{
		repo: repo,
	}
}

func (s *service) AddComment(userID uint, photoID uint, message string) (*domain.Comment, error) {
	comment := &domain.Comment{
		UserID:  userID,
		PhotoID: photoID,
		Message: message,
	}

	return s.repo.SaveComment(comment)
}

func (s *service) GetCommentsByUserID(userID uint) (*[]domain.Comment, error) {
	return s.repo.GetCommentsByUserID(userID)
}

func (s *service) UpdateComment(commentID uint, message string) (*domain.Comment, error) {
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}

	comment.Message = message

	return s.repo.UpdateComment(comment)
}

func (s *service) DeleteComment(commentID uint) error {
	return s.repo.DeleteCommentByID(commentID)
}
