package main

import (
	"github.com/adamnasrudin03/library/config"
	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/repository"
	"github.com/adamnasrudin03/library/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB                 		= config.SetupDbConnection()

	publisherRepo 		repository.PublisherRepository  = repository.NewPublisherRepository(db)

	publisherService 	service.PublisherService 		= service.NewPublisherService(publisherRepo)
)

func main() {
	defer config.CloseDbConnection(db)

	inputPublisher := dto.CreatePublisher{
		Name: "Test service",
		Position: "pustakawan",
		Email: "test@gmail.com",
		Password: "password",
	}
	publisherService.CreatePublisher(inputPublisher)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "Welcome my application",
		})
	})

	router.Run()
}