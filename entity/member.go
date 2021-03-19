package entity

import "time"

//Member represents members table in database
type Member struct {
	ID       		uint64  	`gorm:"primary_key:auto_increment" json:"id"`
	PublisherID     uint64  	`gorm:"not null" json:"-"`
	Name     		string 		`gorm:"type:varchar(255)" json:"name"`
	Profession    	string  	`gorm:"type:varchar(255)" json:"profession"`
	Email    		string  	`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Address    		string  	`gorm:"type:text" json:"address"`
	CreatedAt   	time.Time 	`json:"-"`
	UpdatedAt   	time.Time	`json:"-"`
	Publisher		Publisher	`gorm:"foreignkey:PublisherID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"publisher"`

}
