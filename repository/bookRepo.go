package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type BookRepository interface {
	Save(book entity.Book) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(bookID uint64) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
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

func (r *bookRepository) FindByID(bookID uint64) (entity.Book, error) {
	var book entity.Book

	err := r.db.Preload("Publisher").Where("id = ?", bookID).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Update(book entity.Book) (entity.Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	r.db.Preload("Publisher").Find(&book)

	return book, nil
}
