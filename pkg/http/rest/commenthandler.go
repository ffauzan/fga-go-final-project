package rest

import (
	"errors"
	"final-project/pkg/domain"
	"net/http"
	"strconv"
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

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	// Bind request body to UpdateCommentRequest struct
	// TODO: Add validation
	var req UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Get commentID from path
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Check if comment userID is equal to currentUserID
	comment, err := h.commentService.GetCommentByID(uint(commentID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}
	if comment.UserID != currentUserID {
		SendErrorResponse(c, err, http.StatusUnauthorized)
		return
	}

	// Update comment
	comment, err = h.commentService.UpdateComment(uint(commentID), req.Message)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send response
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func (h *CommentHandler) GetCommentsByUserID(c *gin.Context) {
	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Get comments
	comments, err := h.commentService.GetCommentsByUserID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get user
	user, err := h.userService.GetUserByID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send response
	commentResponses := formatCommentsOfUser(user, comments, h.photoService)

	c.JSON(http.StatusOK, commentResponses)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Get commentID from path
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Check if comment userID is equal to currentUserID
	comment, err := h.commentService.GetCommentByID(uint(commentID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}
	if comment.UserID != currentUserID {
		SendErrorResponse(c, errors.New("insufficient privileges"), http.StatusUnauthorized)
		return
	}

	// Delete comment
	err = h.commentService.DeleteComment(uint(commentID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send response
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Your comment ahs been deleted successfully",
	})
}
