package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FirstName string         `gorm:"not null" json:"first_name" form:"first_name"`
	LastName  string         `gorm:"not null" json:"last_name" form:"last_name"`
	Email     string         `gorm:"not null;unique" json:"email" form:"email"`
	Phone     string         `gorm:"not null;unique" json:"phone" form:"phone"`
	Balance   int            `gorm:"not null" json:"balance" form:"balance"`
	Books     []*Book        `gorm:"foreignKey:Owner;references:ID"`
}
