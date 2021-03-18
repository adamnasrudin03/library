package entity

import "time"

//Publisher represents publishers table in database
type Publisher struct {
	ID       	uint64  	`gorm:"primary_key:auto_increment" json:"id"`
	Name     	string 		`gorm:"type:varchar(255)" json:"name"`
	Position    string  	`gorm:"type:varchar(255)" json:"position"`
	Email    	string  	`gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password 	string  	`gorm:"->;<-;not null" json:"-"`
	Token    	string   	`gorm:"-" json:"token,omitempty"`
	CreatedAt   time.Time 	
	UpdatedAt   time.Time	

}
