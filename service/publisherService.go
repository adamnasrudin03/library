package service

import (
	"errors"

	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
	"golang.org/x/crypto/bcrypt"
)


type PublisherService interface {
	CreatePublisher(input dto.CreatePublisher) (entity.Publisher, error)
	LoginPublisher(input dto.LoginPublisher) (entity.Publisher, error)
	FindByIdPublisher(ID uint64) (entity.Publisher, error)
	UpdatePublisher(input dto.UpdatePublisher) (entity.Publisher, error)
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

func (s *service) LoginPublisher(input dto.LoginPublisher) (entity.Publisher, error) {
	email := input.Email
	password := input.Password

	publisher, err := s.repository.FindByEmail(email)
	if err != nil {
		return publisher, err
	}

	if publisher.ID == 0 {
		return publisher, errors.New("publisher no found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(publisher.Password), []byte(password))
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}

func (s *service) FindByIdPublisher(ID uint64) (entity.Publisher, error) {
	publisher, err := s.repository.FindById(ID)
	if err != nil {
		return publisher, err
	}

	if publisher.ID == 0 {
		return publisher, errors.New("publisher not found on with that ID")
	}

	return publisher, nil
}

func (s *service) UpdatePublisher(input dto.UpdatePublisher) (entity.Publisher, error) {
	publisher, err := s.repository.FindById(input.ID)
	if err != nil {
		return publisher, err
	}

	publisher.Name = input.Name
	publisher.Email = input.Email
	publisher.Password = input.Password

	updatedPublisher, err := s.repository.Update(publisher)
	if err != nil {
		return updatedPublisher, err
	}

	return updatedPublisher, nil
}