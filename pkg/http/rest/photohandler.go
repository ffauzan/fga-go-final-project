package rest

import (
	"final-project/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type PhotoHandler struct {
	photoService domain.PhotoService
}

func NewPhotoHandler(photoService domain.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
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
