package service

import (
	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
)


type MemberService interface {
	CreateMember(input dto.CreateMember) (entity.Member, error)
}

type memberService struct {
	repository repository.MemberRepository
}

func NewMemberService(repository repository.MemberRepository) *memberService {
	return &memberService{repository}
}

func (s *memberService) CreateMember(input dto.CreateMember) (entity.Member, error) {
	member := entity.Member{}
	member.Name = input.Name
	member.Gender = input.Gender
	member.Profession = input.Profession
	member.Email = input.Email
	member.Address = input.Address
	member.PublisherID = input.Publisher.ID

	newMember, err := s.repository.Save(member)
	if err != nil {
		return newMember, err
	}

	return newMember, nil
}