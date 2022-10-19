package sqldb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Order{})

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
