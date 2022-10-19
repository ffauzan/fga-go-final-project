package sqldb

import (
	"final-project/pkg/domain"
	"time"

	"gorm.io/gorm"
)

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

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(user *domain.User) (*domain.User, error) {
	dbUser := User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}

	err := r.db.Create(&dbUser).Error
	if err != nil {
		return nil, err
	}

	user.ID = dbUser.ID
	user.CreatedAt = dbUser.CreatedAt
	user.UpdatedAt = dbUser.UpdatedAt

	return user, nil
}

func (r *UserRepository) GetUserByID(userID uint) (*domain.User, error) {
	var dbUser User
	err := r.db.First(&dbUser, userID).Error
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Age:      dbUser.Age,
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var dbUser User
	err := r.db.Where("username = ?", username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Password: dbUser.Password,
		Email:    dbUser.Email,
		Age:      dbUser.Age,
	}

	return &user, nil
}

func (r *UserRepository) DeleteUserByID(userID uint) error {
	err := r.db.Delete(&User{}, userID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	dbUser := User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	err := r.db.Save(&dbUser).Error
	if err != nil {
		return nil, err
	}

	user.UpdatedAt = dbUser.UpdatedAt

	return user, nil
}

func (r *UserRepository) IsUsernameExist(username string) bool {
	var dbUser User
	err := r.db.Where("username = ?", username).First(&dbUser).Error
	return err == nil
}

func (r *UserRepository) IsEmailExist(email string) bool {
	var dbUser User
	err := r.db.Where("email = ?", email).First(&dbUser).Error
	return err == nil
}

func (r *UserRepository) GetUsers() ([]*domain.User, error) {
	var dbUsers []User
	err := r.db.Find(&dbUsers).Error
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = &domain.User{
			ID:       dbUser.ID,
			Username: dbUser.Username,
			Email:    dbUser.Email,
			Age:      dbUser.Age,
		}
	}

	return users, nil
}
