package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Body  string `gorm:"type:text;not null" json:"body"`
}