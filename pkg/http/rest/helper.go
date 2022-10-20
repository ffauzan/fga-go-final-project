package rest

import "final-project/pkg/domain"

func formatPhotosOfUser(user domain.User, photos []domain.Photo) []PhotoOfUserResponse {
	var photosOfUser []PhotoOfUserResponse
	for _, photo := range photos {
		photosOfUser = append(photosOfUser, PhotoOfUserResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: PhotoUser{
				Email:    user.Email,
				Username: user.Username,
			},
		})
	}
	return photosOfUser
}
