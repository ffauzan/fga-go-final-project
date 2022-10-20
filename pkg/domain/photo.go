package domain

import "time"

type Photo struct {
	ID        uint
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}

type AddPhotoRequest struct {
	Title    string
	Caption  string
	PhotoUrl string
}

type PhotoService interface {
	SavePhoto(userID uint, req *AddPhotoRequest) (*Photo, error)
	GetPhotoByID(photoID uint) (*Photo, error)
	GetPhotosByUserID(userID uint) (*[]Photo, error)
	UpdatePhoto(photoID uint, req *AddPhotoRequest) (*Photo, error)
	DeletePhoto(photoID uint) error
}

type PhotoRepository interface {
	SavePhoto(photo *Photo) (*Photo, error)
	GetPhotoByID(photoID uint) (*Photo, error)
	UpdatePhoto(photo *Photo) (*Photo, error)
	GetPhotosByUserID(userID uint) (*[]Photo, error)
	DeletePhotoByID(photoID uint) error
}
