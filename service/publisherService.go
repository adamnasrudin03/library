package service

import (
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
)


type PublisherService interface {
	CreateCampaign() (entity.Publisher, error)
}

type service struct {
	repository repository.PublisherRepository
}

func NewPublisherService(repository repository.PublisherRepository) *service {
	return &service{repository}
}