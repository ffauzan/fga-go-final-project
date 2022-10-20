package rest

import (
	"final-project/pkg/domain"
	"net/http"
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
