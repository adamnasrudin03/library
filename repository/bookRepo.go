package repository

import "gorm.io/gorm"

type BookRepository interface {
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}