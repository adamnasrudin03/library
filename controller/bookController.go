package controller

import (
	"net/http"

	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/helper"
	"github.com/adamnasrudin03/library/service"
	"github.com/gin-gonic/gin"
)



type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *bookController{
	return &bookController{bookService}
}

func (c *bookController) CreateBook(ctx *gin.Context) {
	var input dto.CreateBook

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create book", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentPublisher := ctx.MustGet("currentPublisher").(entity.Publisher)
	input.Publisher = currentPublisher

	newBook, err := c.bookService.CreateBook(input)
	if err != nil {
		response := helper.APIResponse("Failed to create Book", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create Book", http.StatusOK, "success", newBook)
	ctx.JSON(http.StatusOK, response)

}