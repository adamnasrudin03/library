package main

import (
	"net/http"

	"github.com/adamnasrudin03/library/config"
	"github.com/adamnasrudin03/library/controller"
	"github.com/adamnasrudin03/library/middleware"
	"github.com/adamnasrudin03/library/repository"
	"github.com/adamnasrudin03/library/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB                 		= config.SetupDbConnection()

	publisherRepo 		repository.PublisherRepository  = repository.NewPublisherRepository(db)
	memberRepo 			repository.MemberRepository 	= repository.NewMemberRepository(db)
	bookRepo			repository.BookRepository		= repository.NewBookRepository(db)

	publisherService 	service.PublisherService 		= service.NewPublisherService(publisherRepo)
	authService			service.AuthService				= service.NewAuthService(publisherRepo)
	memberService 		service.MemberService 			= service.NewMemberService(memberRepo)
	bookService			service.BookService				= service.NewBookService(bookRepo)
)

func main() {
	defer config.CloseDbConnection(db)

	authMiddleware := middleware.NewAuthMiddleware(authService, publisherService)

	publisherController :=	controller.NewPublisherController(publisherService, authService)
	memberController := controller.NewMemberController(memberService)
	bookController := controller.NewBookController(bookService)

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
	api.POST("/auth/sessions", publisherController.Login)
	api.PUT("/publishers",authMiddleware.AuthorizationMiddleware() ,publisherController.Update)

	api.POST("/members", authMiddleware.AuthorizationMiddleware(), memberController.CreateMember)
	api.GET("/members", authMiddleware.AuthorizationMiddleware(), memberController.FindAllMember)
	api.GET("/members/:id", authMiddleware.AuthorizationMiddleware(), memberController.FindByIDMember)
	api.PUT("/members/:id/update", authMiddleware.AuthorizationMiddleware(), memberController.UpdateMember)
	api.DELETE("/members/:id/delete", authMiddleware.AuthorizationMiddleware(), memberController.DeleteByIDMember)

	api.POST("/books", authMiddleware.AuthorizationMiddleware(), bookController.CreateBook)
	api.GET("/books", authMiddleware.AuthorizationMiddleware(), bookController.FindAllBook)
	api.GET("/books/:id", authMiddleware.AuthorizationMiddleware(), bookController.FindByIDBook)
	api.PUT("/books/:id/update", authMiddleware.AuthorizationMiddleware(), bookController.UpdateBook)
	api.DELETE("/books/:id/delete", authMiddleware.AuthorizationMiddleware(), bookController.DeleteByIDBook)

	router.NoRoute(func(c *gin.Context) {
    	c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
			"code": http.StatusNotFound,
			"status": "error",
		})
	})

	router.Run()
}