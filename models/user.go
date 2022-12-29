package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name" form:"first_name"`
	LastName  string `gorm:"not null" json:"last_name" form:"last_name"`
	Email     string `gorm:"not null;unique" json:"email" form:"email"`
	Phone     string `gorm:"not null;unique" json:"phone" form:"phone"`
	Balance   int    `gorm:"not null" json:"balance" form:"balance"`
	Books     []Book `gorm:"foreignKey:Owner;references:ID"`
}
