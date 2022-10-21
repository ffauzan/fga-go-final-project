package rest

import (
	"errors"
	"final-project/pkg/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AddSocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type SocialMediaOfUserResponse struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	SocialMediaUrl string          `json:"social_media_url"`
	UserID         uint            `json:"user_id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	User           SocialMediaUser `json:"User"`
}

type SocialMediaUser struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type SocialMediaHandler struct {
	SocialMediaService domain.SocialMediaService
	UserService        domain.UserService
}

func NewSocialMediaHandler(socialMediaService domain.SocialMediaService, userService domain.UserService) *SocialMediaHandler {
	return &SocialMediaHandler{
		SocialMediaService: socialMediaService,
		UserService:        userService,
	}
}

func (h *SocialMediaHandler) AddSocialMedia(c *gin.Context) {
	// Bind the request body to the AddSocialMediaRequest struct
	var req AddSocialMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get the user ID from the request context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Save the social media
	socialMedia, err := h.SocialMediaService.AddSocialMedia(currentUserID, req.Name, req.SocialMediaUrl)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send the response
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"createdAt":        socialMedia.CreatedAt,
	})
}

func (h *SocialMediaHandler) GetSocialMedias(c *gin.Context) {
	// Get the user ID from the request context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Get user
	user, err := h.UserService.GetUserByID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get social medias
	socialMedias, err := h.SocialMediaService.GetSocialMediasByUserID(currentUserID)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	res := formatSocialMediaOfUser(user, socialMedias)

	// Send the response
	c.JSON(http.StatusOK, map[string]interface{}{
		"social_medias": res,
	})
}

func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	// Bind the request body to the AddSocialMediaRequest struct
	var req AddSocialMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get socialMediaID from URL
	socialMediaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Get the user ID from the request context
	currentUserID := c.MustGet("currentUserID").(uint)

	// Check if social media userID is equal to current userID
	socialMedia, err := h.SocialMediaService.GetSocialMediaByID(uint(socialMediaID))
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	if socialMedia.UserID != currentUserID {
		SendErrorResponse(c, errors.New("insufficient privileges"), http.StatusUnauthorized)
		return
	}

	// Update the social media
	socialMedia, err = h.SocialMediaService.UpdateSocialMedia(uint(socialMediaID), req.Name, req.SocialMediaUrl)
	if err != nil {
		SendErrorResponse(c, err, http.StatusBadRequest)
		return
	}

	// Send the response
	c.JSON(http.StatusOK, map[string]interface{}{
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
	})
}
