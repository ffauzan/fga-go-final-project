package domain

import "time"

type Comment struct {
	ID        uint
	UserID    uint
	PhotoID   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
