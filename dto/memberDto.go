package dto

import "github.com/adamnasrudin03/library/entity"

//CreateMember is used when client post
type CreateMember struct {
	Name     		string 		`json:"name" form:"name" binding:"required"`
	Gender     		string 		`json:"gender" form:"gender" binding:"required"`
	Profession    	string  	`json:"profession" form:"profession" binding:"required"`
	Email    		string  	`json:"email" form:"email" binding:"required,email" `
	Address    		string  	`json:"address" form:"address" `
	Publisher		entity.Publisher
}