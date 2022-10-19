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
