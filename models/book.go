package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Owner     int
	Name      string
	Author    string
	Publisher string
	ISBN      string
	Price     string
	Stock     int
}
