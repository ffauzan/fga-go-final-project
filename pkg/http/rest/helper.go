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

func formatCommentsOfUser(user *domain.User, comments *[]domain.Comment, photoService domain.PhotoService) []CommentOfUserResponse {
	var commentsOfUser []CommentOfUserResponse
	for _, comment := range *comments {
		// Get photo
		photo, err := photoService.GetPhotoByID(comment.PhotoID)
		if err != nil {
			return nil
		}

		commentsOfUser = append(commentsOfUser, CommentOfUserResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			User: CommentUser{
				ID:       user.ID,
				Email:    user.Email,
				Username: user.Username,
			},
			Photo: CommentPhoto{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoUrl: photo.PhotoUrl,
				UserID:   photo.UserID,
			},
		})
	}
	return commentsOfUser
}
