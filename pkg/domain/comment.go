package domain

import "time"

type Comment struct {
	ID        uint
	UserID    uint
	PhotoID   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentService interface {
	AddComment(userID uint, photoID uint, message string) (*Comment, error)
	GetCommentsByUserID(userID uint) (*[]Comment, error)
	UpdateComment(commentID uint, message string) (*Comment, error)
	DeleteComment(commentID uint) error
}

type CommentRepository interface {
	SaveComment(comment *Comment) (*Comment, error)
	GetCommentByID(commentID uint) (*Comment, error)
	GetCommentsByUserID(userID uint) (*[]Comment, error)
	UpdateComment(comment *Comment) (*Comment, error)
	DeleteCommentByID(commentID uint) error
}
