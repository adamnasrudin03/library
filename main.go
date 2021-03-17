package main

import (
	"github.com/adamnasrudin03/library/config"
	"github.com/gin-gonic/gin"
)

var (
	db             *gorm.DB                  = config.SetupDbConnection()
)

func main() {
	defer config.CloseDbConnection(db)

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