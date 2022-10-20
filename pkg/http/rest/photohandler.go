package rest

import (
	"errors"
	"final-project/pkg/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AddPhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type PhotoHandler struct {
	photoService domain.PhotoService
	userService  domain.UserService
}

type PhotoOfUserResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      PhotoUser `json:"User"`
}

type PhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewPhotoHandler(photoService domain.PhotoService, userService domain.UserService) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
		userService:  userService,
	}
}

func (h *PhotoHandler) AddPhoto(c *gin.Context) {
	// Bind request body to AddPhotoRequest struct
	// TODO: Add validation
	var req AddPhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	photo, err := h.photoService.SavePhoto(currentUserID, &domain.AddPhotoRequest{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	})

	if err != nil {
		SendErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {
	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Get photos of current user
	photos, err := h.photoService.GetPhotosByUserID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	// Get the user corressponding to the userID
	user, err := h.userService.GetUserByID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	// Format json response
	photosOfUserResponse := formatPhotosOfUser(*user, *photos)

	c.JSON(http.StatusOK, photosOfUserResponse)
}

func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {
	// Bind request body to AddPhotoRequest struct
	var req AddPhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get photoID from URL
	photoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get currentUserID from context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Check if photo userID equal to current userID
	photo, err := h.photoService.GetPhotoByID(uint(photoID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	if photo.UserID != currentUserID {
		SendErrorResponse(c, errors.New("you don't have access to this photo"), http.StatusUnauthorized)
		return
	}

	// Update photo
	photo, err = h.photoService.UpdatePhoto(uint(photoID), &domain.AddPhotoRequest{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	})

	if err != nil {
		SendErrorResponse(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}