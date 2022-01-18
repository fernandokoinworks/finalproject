package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	UserID    int    `json:"userId" example:"1"`
	Title     string `json:"title" example:"Send email to Leo messi"`
	Completed bool   `json:"completed" example:"true"`
}

type ReqTodo struct {
	UserID    int    `json:"userId" example:"1"`
	Title     string `json:"title" example:"Send email to Leo messi"`
	Completed bool   `json:"completed" example:"true"`
}
