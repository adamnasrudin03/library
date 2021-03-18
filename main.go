package main

import (
	"github.com/adamnasrudin03/library/config"
	"github.com/adamnasrudin03/library/controller"
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
	authService			service.AuthService				= service.NewAuthService()
)

func main() {
	defer config.CloseDbConnection(db)
	publisherController :=	controller.NewPublisherController(publisherService, authService)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"message": "Welcome my application",
		})
	})

	api := router.Group("/api/v1")

	api.POST("/auth/publishers", publisherController.RegisterPublisher)

	router.Run()
}