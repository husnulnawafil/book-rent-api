package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name" form:"first_name"`
	Email     string `gorm:"not null" json:"email" form:"email"`
	Phone     string `gorm:"not null" json:"phone" form:"phone"`
	LastName  string `gorm:"not null" json:"last_name" form:"last_name"`
	Balance   int    `gorm:"not null" json:"balance" form:"balance"`
}
