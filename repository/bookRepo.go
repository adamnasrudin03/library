package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type BookRepository interface {
	Save(book entity.Book) (entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Save(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}
	r.db.Preload("Publisher").Find(&book)

	return book, nil
}
