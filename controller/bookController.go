package controller

import (
	"net/http"
	"strconv"

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

func (c *bookController) FindAllBook(ctx *gin.Context) {
	books, err := c.bookService.FindAllBook()
	if err != nil {
		response := helper.APIResponse("Error to get books", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of books", http.StatusOK, "success", books)
	ctx.JSON(http.StatusOK, response)
}

func (c *bookController) FindByIDBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponseError("Param id not found / did not match", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book, err := c.bookService.FindByIDBook(id)
	if err != nil {
		response := helper.APIResponse("Error to get book", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if (book == entity.Book{}) {
		response := helper.APIResponse("Book not found", http.StatusNotFound, "success", nil)
		ctx.JSON(http.StatusNotFound, response)
	} else {
		response := helper.APIResponse("List of Detail book", http.StatusOK, "success", book)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *bookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.APIResponseError("Param id not found / did not match", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book, _ := c.bookService.FindByIDBook(id)
	if (book == entity.Book{}) {
		response := helper.APIResponse("Book not found", http.StatusNotFound, "success", nil)
		ctx.JSON(http.StatusNotFound, response)
	}

	var input dto.UpdateBook
	if input.InitialStock == 0 {
		input.InitialStock = book.InitialStock
	}

	currentPublisher := ctx.MustGet("currentPublisher").(entity.Publisher)
	input.Publisher = currentPublisher

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update book", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateBook, err := c.bookService.UpdateBook(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to updated book", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated book", http.StatusOK, "success", updateBook)
	ctx.JSON(http.StatusOK, response)
}
