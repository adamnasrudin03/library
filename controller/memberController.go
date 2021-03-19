package controller

import "github.com/adamnasrudin03/library/service"

type memberController struct {
	memberService service.MemberService
}

func NewMemberController(memberService service.MemberService) *memberController {
	return &memberController{memberService}
}