package domain

import "time"

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

type UserService interface {
	DeleteUser(userID uint) error
	UpdateUser(userID uint, req *UpdateUserRequest) (*User, error)
	IsUserExist(userID uint) bool
	Register(req *RegisterRequest) (*User, error)
	Login(req *LoginRequest) (*string, error)
}

type UserRepository interface {
	SaveUser(user *User) (*User, error)
	GetUserByID(userID uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
	DeleteUserByID(userID uint) error
	UpdateUser(userID *User) (*User, error)
	IsUsernameExist(username string) bool
	IsEmailExist(email string) bool
}
