package main

import (
	"fmt"

	"github.com/adamnasrudin03/library/config"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                 		= config.SetupDbConnection()
	publisherRepo 	repository.PublisherRepository  = repository.NewPublisherRepository(db)
)

func main() {
	defer config.CloseDbConnection(db)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error " ,err)
	}
	publisher := entity.Publisher{
		Name: "ADAM nasrudin",
		Email: "adam@gmail.com",
		Password: string(passwordHash),
	}
	publisherRepo.Save(publisher) 

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