package dto

//CreatePublisher is used when client post
type CreatePublisher struct {
	Name     	string `json:"name" form:"name" binding:"required"`
	Position    string `json:"position" form:"position" binding:"required"`
	Email    	string `json:"email" form:"email" binding:"required,email" `
	Password 	string `json:"password" form:"password" binding:"required"`
}

//LoginPublisher is used when client post
type LoginPublisher struct {
	Email    	string `json:"email" form:"email" binding:"required,email" `
	Password 	string `json:"password" form:"password" binding:"required"`
}

//UpdatePublisher is used when client post
type UpdatePublisher struct {
	ID      	uint64 `json:"id" form:"id" binding:"required"`
	Name     	string `json:"name" form:"name" `
	Position    string `json:"position" form:"position" `
	Email    	string `json:"email" form:"email" binding:"required,email" `
	Password 	string `json:"password,omitempty" form:"password,omitempty" binding:"required" `
}