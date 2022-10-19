package domain

import "time"

type SocialMedia struct {
	ID             uint
	Name           string
	SocialMediaUrl string
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
