package database

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
