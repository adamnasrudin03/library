package controller

import "github.com/adamnasrudin03/library/service"

type publisherController struct {
	publisherService service.PublisherService
}

func NewPublisherController(publisherService service.PublisherService) *publisherController {
	return &publisherController{publisherService}
}