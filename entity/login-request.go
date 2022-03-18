package entity

type LoginRequest struct {
	Username string `json:"Username" form:"Username" binding:"required"`
	Password string `json:"Password" form:"Password" binding:"required"`
	Roles    []int  `json:"Roles" form:"Roles" binding:"required"`
}
