package main

import "github.com/gin-gonic/gin"


func main() {
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