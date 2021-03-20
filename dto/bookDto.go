package dto

import "github.com/adamnasrudin03/library/entity"

//CreateBook is used when client post
type CreateBook struct {
	Name     		string 		`json:"name" form:"name" binding:"required"`
	Author     		string 		`json:"author" form:"author" binding:"required"`
	InitialStock    uint32  	`json:"initial_stock" form:"initial_stock" binding:"required"`
	Publisher		entity.Publisher
}

