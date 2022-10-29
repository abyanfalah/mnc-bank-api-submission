package model

type Customer struct {
	Id       string `json:"id" form:"id" db:"id" `
	Name     string `json:"name" form:"name" db:"name" binding:"required"`
	Username string `json:"username" form:"username" db:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password" db:"password" binding:"required"`
	Balance  int    `json:"balance" form:"balance" db:"balance" binding:"required"`
}

type Credential struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type Session struct {
	Id       string `json:"id" form:"id" db:"id" `
	Username string `json:"username" db:"username" binding:"required"`
	Name     string `json:"name" form:"name" db:"name" binding:"required"`
}
