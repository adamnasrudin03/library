package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type MemberRepository interface {
	Save(member entity.Member) (entity.Member, error)
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

	return member, nil
}