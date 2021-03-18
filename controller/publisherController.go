package controller

import (
	"fmt"
	"net/http"

	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/helper"
	"github.com/adamnasrudin03/library/service"
	"github.com/gin-gonic/gin"
)


type publisherController struct {
	publisherService service.PublisherService
	authService service.AuthService
}

func NewPublisherController(publisherService service.PublisherService, authService service.AuthService) *publisherController {
	return &publisherController{publisherService, authService}
}

func (c *publisherController) RegisterPublisher(ctx *gin.Context){
	var input dto.CreatePublisher
	
	err := ctx.BindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	if !c.authService.IsDuplicateEmail(input.Email) {
		errorMessage := fmt.Sprintf("email has been registered \nDuplicate email : %s ", input.Email)
		response := helper.APIResponseError("Register account failed", http.StatusConflict, "error", errorMessage)
		ctx.JSON(http.StatusConflict, response)
		return
	}

	publisher, err := c.publisherService.CreatePublisher(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := c.authService.GenerateToken(publisher.ID, publisher.Name)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	publisher.Token = token

	response := helper.APIResponse("Account has ben registered", http.StatusOK, "success", publisher)
	ctx.JSON(http.StatusOK, response)

}

func (c *publisherController) Login(ctx *gin.Context) {
	var input dto.LoginPublisher

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	publisher, err := c.publisherService.LoginPublisher(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := c.authService.GenerateToken(publisher.ID, publisher.Name)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	publisher.Token = token
	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", publisher)
	ctx.JSON(http.StatusOK, response)

}

func (c *publisherController) Update(ctx *gin.Context) {
	currentPublisher := ctx.MustGet("currentPublisher").(entity.Publisher)

	var input dto.UpdatePublisher
	input.ID = currentPublisher.ID

	if input.Name == "" {
		input.Name = currentPublisher.Name
	}
	if input.Position == "" {
		input.Position = currentPublisher.Position
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedPublisher, err := c.publisherService.UpdatePublisher(input)
	if err != nil {
		response := helper.APIResponse("Failed to updated campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to updated campaign", http.StatusOK, "success", updatedPublisher)
	ctx.JSON(http.StatusOK, response)
}