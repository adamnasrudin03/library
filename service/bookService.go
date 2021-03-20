package service

import (
	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
)


type BookService interface {
	CreateBook(input dto.CreateBook) (entity.Book, error)
	FindAllBook() ([]entity.Book, error)
	FindByIDBook(bookID uint64) (entity.Book, error)
	UpdateBook(bookID uint64, input dto.UpdateBook) (entity.Book, error)
	DeleteByIDBook(bookID uint64) (entity.Book, error)
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) CreateBook(input dto.CreateBook) (entity.Book, error) {
	book := entity.Book{}
	book.Name = input.Name
	book.Author = input.Author
	book.InitialStock = input.InitialStock
	book.CurrentStock = input.InitialStock
	book.TotalBorrowers = 0
	book.Avaliable = 1
	book.PublisherID = input.Publisher.ID

	newBook, err := s.repository.Save(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *bookService) FindAllBook() ([]entity.Book, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}
	
	return books, nil
}

func (s *bookService) FindByIDBook(bookID uint64) (entity.Book, error) {
	book, err := s.repository.FindByID(bookID)
	if err != nil {
		return book, err
	}
	
	return book, nil
}

func (s *bookService) UpdateBook(bookID uint64, input dto.UpdateBook) (entity.Book, error) {
	book, err := s.repository.FindByID(bookID)
	if err != nil {
		return book, err
	}
	
	book.Name = input.Name
	book.Author = input.Author
	book.InitialStock = input.InitialStock
	book.CurrentStock = input.CurrentStock
	book.Avaliable = input.Avaliable
	book.Publisher.ID = input.Publisher.ID

	newBook, err := s.repository.Update(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *bookService) DeleteByIDBook(bookID uint64) (entity.Book, error) {
	book, err := s.repository.DeleteByID(bookID)
	if err != nil {
		return book, err
	}
	
	return book, nil
}
