package rest

import (
	"final-project/pkg/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddCommentRequest struct {
	Message string `json:"message"`
	PhotoID uint   `json:"photo_id"`
}

type UpdateCommentRequest struct {
	Message string `json:"message"`
}

type CommentOfUserResponse struct {
	ID        uint         `json:"id"`
	Message   string       `json:"message"`
	PhotoID   uint         `json:"photo_id"`
	UserID    uint         `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      CommentUser  `json:"User"`
	Photo     CommentPhoto `json:"Photo"`
}

type CommentUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CommentPhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type CommentHandler struct {
	commentService domain.CommentService
	userService    domain.UserService
	photoService   domain.PhotoService
}

func NewCommentHandler(commentService domain.CommentService, userService domain.UserService, photoService domain.PhotoService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		userService:    userService,
		photoService:   photoService,
	}
}

func (h *CommentHandler) AddComment(c *gin.Context) {
	// Bind request body to AddCommentRequest struct
	// TODO: Add validation
	var req AddCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Save comment
	comment, err := h.commentService.AddComment(currentUserID, req.PhotoID, req.Message)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send response
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}
