package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Author string `json:"Author"`
}
