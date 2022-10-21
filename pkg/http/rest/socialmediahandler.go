package rest

import (
	"final-project/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddSocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
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
