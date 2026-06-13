package repository

import "gorm.io/gorm"

type PostRepository struct {
	db *gorm.DB
}

