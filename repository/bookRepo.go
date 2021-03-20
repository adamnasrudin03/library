package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type BookRepository interface {
	Save(book entity.Book) (entity.Book, error)
	FindAll() ([]entity.Book, error)
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

func (r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Preload("Publisher").Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}
