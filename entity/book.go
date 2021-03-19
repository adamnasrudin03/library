package entity

import "time"

//Book represents books table in database
type Book struct {
	ID       		uint64  	`gorm:"primary_key:auto_increment" json:"id"`
	PublisherID     uint64  	`json:"-"`
	Name     		string 		`gorm:"type:varchar(255)" json:"name"`
	Author     		string 		`gorm:"type:varchar(255)" json:"author"`
	InitialStock    uint32  	`gorm:"type:int(11)" json:"initial_stock"`
	CurrentStock   	uint32  	`gorm:"type:int(11)" json:"current_stock"`
	TotalBorrowers  uint32  	`gorm:"type:int(11)" json:"total_borrowers"`
	Avaliable    	int		  	`gorm:"type:tinyint(4)" json:"avaliable"`
	CreatedAt   	time.Time 	`json:"-"`
	UpdatedAt   	time.Time	`json:"-"`
	Publisher		Publisher	`gorm:"foreignkey:PublisherID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"publisher"`
}