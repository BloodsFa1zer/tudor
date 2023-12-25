package models

type InLogin struct {
	Email    string `json:"email"      form:"email"		 binding:"required"`
	Password string `json:"password"   form:"password"	 binding:"required"`
}

type InUserInfo struct {
	Id int64 `uri:"id" binding:"required"`
}

type EmailRequest struct {
	Email string `json:"email"	form:"email" 	binding:"required,email"`
}

type Id struct {
	Id int64 `json:"id" binding:"required"`
}

type UserPassword struct {
	Password string `json:"password" form:"password" binding:"required"`
}
