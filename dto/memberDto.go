package dto

//CreateMember is used when client post
type CreateMember struct {
	PublisherID     uint64  	`json:"publisher_id" form:"publisher_id" binding:"required"`
	Name     		string 		`json:"name" form:"name" binding:"required"`
	Gender     		string 		`json:"gender" form:"gender" binding:"required"`
	Profession    	string  	`json:"profession" form:"profession" binding:"required"`
	Email    		string  	`json:"email" form:"email" binding:"required,email" `
	Address    		string  	`json:"address" form:"address" `
}