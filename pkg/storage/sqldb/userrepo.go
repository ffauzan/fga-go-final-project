package sqldb

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null;unique"`
	Email        string `gorm:"not null;unique"`
	Password     string `gorm:"not null"`
	Age          int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo       `gorm:"foreignKey:UserID"`
	Comments     []Comment     `gorm:"foreignKey:UserID"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID"`
}

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
