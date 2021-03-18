package service

import (
	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
	"golang.org/x/crypto/bcrypt"
)


type PublisherService interface {
	CreatePublisher(input dto.CreatePublisher) (entity.Publisher, error)
}

type service struct {
	repository repository.PublisherRepository
}

func NewPublisherService(repository repository.PublisherRepository) *service {
	return &service{repository}
}

func (s *service) CreatePublisher(input dto.CreatePublisher) (entity.Publisher, error) {
	publisher := entity.Publisher{}
	publisher.Name = input.Name
	publisher.Position = input.Position
	publisher.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return publisher, err
	}

	publisher.Password = string(passwordHash)

	newPublisher, err := s.repository.Save(publisher)
	if err != nil {
		return newPublisher, err
	}

	return newPublisher, nil
}