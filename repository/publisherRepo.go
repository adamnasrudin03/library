package repository

import (
	"github.com/adamnasrudin03/library/entity"
	"gorm.io/gorm"
)


type PublisherRepository interface {
	Save(publisher entity.Publisher) (entity.Publisher, error)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) (entity.Publisher, error)
}

type repository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var publisher entity.Publisher
	return r.db.Where("email = ?", email).Take(&publisher)
}

func (r *repository) Save(publisher entity.Publisher) (entity.Publisher, error) {
	err := r.db.Create(&publisher).Error
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}

func (r *repository) FindByEmail(email string) (entity.Publisher, error) {
	var publisher entity.Publisher

	err := r.db.Where("email = ?", email).Find(&publisher).Error
	if err != nil {
		return publisher, err
	}

	return publisher, nil
}