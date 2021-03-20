package dto

import "github.com/adamnasrudin03/library/entity"

//CreateBook is used when client post
type CreateBook struct {
	Name     		string 		`json:"name" form:"name" binding:"required"`
	Author     		string 		`json:"author" form:"author" binding:"required"`
	InitialStock    uint32  	`json:"initial_stock" form:"initial_stock" binding:"required"`
	Publisher		entity.Publisher
}


//UpdateBook is used when client post
type UpdateBook struct {
	Name     		string 		`json:"name" form:"name" binding:"required"`
	Author     		string 		`json:"author" form:"author" binding:"required"`
	InitialStock    uint32  	`json:"initial_stock" form:"initial_stock" binding:"required"`
	CurrentStock   	uint32  	`json:"current_stock" form:"current_stock" binding:"required"`
	TotalBorrowers  uint32  	`json:"total_borrowers" form:"total_borrowers" binding:"required"`
	Avaliable    	int		  	`json:"avaliable" form:"avaliable" binding:"required"`
	Publisher		entity.Publisher
}
