package controller

import (
	"net/http"

	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/helper"
	"github.com/adamnasrudin03/library/service"
	"github.com/gin-gonic/gin"
)


type memberController struct {
	memberService service.MemberService
}

func NewMemberController(memberService service.MemberService) *memberController {
	return &memberController{memberService}
}

func (c *memberController) CreateMember(ctx *gin.Context) {
	var input dto.CreateMember

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create member", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentPublisher := ctx.MustGet("currentPublisher").(entity.Publisher)
	input.Publisher = currentPublisher

	newMember, err := c.memberService.CreateMember(input)
	if err != nil {
		response := helper.APIResponse("Failed to create Member", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create Member", http.StatusOK, "success", newMember)
	ctx.JSON(http.StatusOK, response)

}

func (c *memberController) FindAllMember(ctx *gin.Context) {
	members, err := c.memberService.FindAllMember()
	if err != nil {
		response := helper.APIResponse("Error to get members", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of members", http.StatusOK, "success", members)
	ctx.JSON(http.StatusOK, response)
}
