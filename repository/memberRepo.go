package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type MemberRepository interface {
	Save(member entity.Member) (entity.Member, error)
	FindAll() ([]entity.Member, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) Save(member entity.Member) (entity.Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	}
	r.db.Preload("Publisher").Find(&member)

	return member, nil
}

func (r *memberRepository) FindAll() ([]entity.Member, error) {
	var members []entity.Member

	err := r.db.Preload("Publisher").Find(&members).Error
	if err != nil {
		return members, err
	}

	return members, nil
}