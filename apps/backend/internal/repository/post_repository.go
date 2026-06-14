package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}