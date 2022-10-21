package sqldb

import (
	"final-project/pkg/domain"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null;unique;type:varchar(255)"`
	Email        string `gorm:"not null;unique;type:varchar(255)"`
	Password     string `gorm:"not null"`
	Age          int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo       `gorm:"foreignKey:UserID"`
	Comments     []Comment     `gorm:"foreignKey:UserID"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	log.Println("UserRepository created")
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
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Age:       dbUser.Age,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
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
	// Transaction to delete user and all of his photos, comments, and social medias
	tx := r.db.Begin()

	// Delete social medias of user
	err := tx.Where("user_id = ?", userID).Delete(&SocialMedia{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete comments of user
	err = tx.Where("user_id = ?", userID).Delete(&Comment{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete photos of user and all of its comments
	// Get all photos of user
	var photos []Photo
	err = tx.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete all comments of photos
	for _, photo := range photos {
		err = tx.Where("photo_id = ?", photo.ID).Delete(&Comment{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Delete photos of user
	err = tx.Where("user_id = ?", userID).Delete(&Photo{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete user
	err = tx.Where("id = ?", userID).Delete(&User{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *UserRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	err := r.db.Model(&User{}).Where("id = ?", user.ID).Updates(User{
		Username:  user.Username,
		Email:     user.Email,
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}

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
