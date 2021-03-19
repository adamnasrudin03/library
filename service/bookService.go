package service

import "github.com/adamnasrudin03/library/repository"

type BookService interface {
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}
