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

type AuthService interface {
	// GenerateToken(user *User) (*string, error)
	GenerateToken(userID uint) (string, error)
	// IsTokenValid(token string) (bool, error)
	// GetUserIDFromToken(token string) (uint, error)
	// IsUserCanCreate(userID uint, entity *interface{}) bool
	// IsUserCanAccess(userID uint, entity *interface{}) bool
	// IsUserCanUpdate(userID uint, entity *interface{}) bool
	// IsUserCanDelete(userID uint, entity *interface{}) bool
}

type CryptoService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(plaintext string, hashed string) error
}
