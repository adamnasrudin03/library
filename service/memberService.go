package service

import (
	"github.com/adamnasrudin03/library/dto"
	"github.com/adamnasrudin03/library/entity"
	"github.com/adamnasrudin03/library/repository"
)


type MemberService interface {
	CreateMember(input dto.CreateMember) (entity.Member, error)
	FindAllMember() ([]entity.Member, error)
	FindByIDMember(memberID uint64) (entity.Member, error)
	UpdateMember(memberID uint64, input dto.UpdateMember)  (entity.Member, error)
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

func (s *memberService) FindAllMember() ([]entity.Member, error) {
	members, err := s.repository.FindAll()
	if err != nil {
		return members, err
	}
	
	return members, nil
}

func (s *memberService) FindByIDMember(memberID uint64) (entity.Member, error) {
	member, err := s.repository.FindByID(memberID)
	if err != nil {
		return member, err
	}
	
	return member, nil
}

func (s *memberService)	UpdateMember(memberID uint64, input dto.UpdateMember) (entity.Member, error){
	member, err := s.repository.FindByID(memberID)
	if err != nil {
		return member, err
	}

	member.Name = input.Name
	member.Gender = input.Gender
	member.Profession = input.Profession
	member.Email = input.Email
	member.Address = input.Address

	newMember, err := s.repository.Update(member)
	if err != nil {
		return newMember, err
	}

	return newMember, nil
}
