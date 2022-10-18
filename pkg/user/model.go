package user

import (
	"time"
)

type User struct {
	ID           uint
	Username     string
	Email        string
	Password     string
	Age          int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

type SocialMedia struct {
	ID             uint
	Name           string
	SocialMediaUrl string
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

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

type Comment struct {
	ID        uint
	UserID    uint
	PhotoID   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginRequest struct {
	Username string
	Password string
}

type RegisterRequest struct {
	Username string
	Email    string
	Age      int
	Password string
}

type UpdateUserRequest struct {
	Username string
	Email    string
}
