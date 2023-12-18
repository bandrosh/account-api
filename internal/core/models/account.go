package models

type Account struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
