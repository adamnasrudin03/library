package service

import (
	"github.com/adamnasrudin03/library/repository"
)


type MemberService interface {
	
}

type memberService struct {
	repository repository.MemberRepository
}

func NewMemberService(repository repository.MemberRepository) *memberService {
	return &memberService{repository}
}
