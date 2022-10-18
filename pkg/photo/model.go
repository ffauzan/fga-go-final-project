package photo

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

type Comment struct {
	ID        uint
	UserID    uint
	PhotoID   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
