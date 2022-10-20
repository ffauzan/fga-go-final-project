package sqldb

import (
	"final-project/pkg/domain"
	"log"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	PhotoID   uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	log.Println("CommentRepository created")
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) SaveComment(comment *domain.Comment) (*domain.Comment, error) {
	dbComment := Comment{
		UserID:  comment.UserID,
		PhotoID: comment.PhotoID,
		Message: comment.Message,
	}

	err := r.db.Create(&dbComment).Error
	if err != nil {
		return nil, err
	}

	comment.ID = dbComment.ID
	comment.CreatedAt = dbComment.CreatedAt
	comment.UpdatedAt = dbComment.UpdatedAt

	return comment, nil
}

func (r *CommentRepository) GetCommentByID(commentID uint) (*domain.Comment, error) {
	var dbComment Comment
	err := r.db.First(&dbComment, commentID).Error
	if err != nil {
		return nil, err
	}

	comment := domain.Comment{
		ID:        dbComment.ID,
		UserID:    dbComment.UserID,
		PhotoID:   dbComment.PhotoID,
		Message:   dbComment.Message,
		CreatedAt: dbComment.CreatedAt,
		UpdatedAt: dbComment.UpdatedAt,
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentsByUserID(userID uint) (*[]domain.Comment, error) {
	var dbComments []Comment
	err := r.db.Where("user_id = ?", userID).Find(&dbComments).Error
	if err != nil {
		return nil, err
	}

	comments := make([]domain.Comment, len(dbComments))
	for i, dbComment := range dbComments {
		comments[i] = domain.Comment{
			ID:        dbComment.ID,
			UserID:    dbComment.UserID,
			PhotoID:   dbComment.PhotoID,
			Message:   dbComment.Message,
			CreatedAt: dbComment.CreatedAt,
			UpdatedAt: dbComment.UpdatedAt,
		}
	}

	return &comments, nil
}

func (r *CommentRepository) UpdateComment(comment *domain.Comment) (*domain.Comment, error) {
	err := r.db.Model(&Comment{}).Where("id = ?", comment.ID).Updates(Comment{
		Message: comment.Message,
	}).Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) DeleteCommentByID(commentID uint) error {
	err := r.db.Delete(&Comment{}, commentID).Error
	if err != nil {
		return err
	}

	return nil
}
