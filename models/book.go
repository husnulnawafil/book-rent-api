package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Owner     uint   `gorm:"not null" json:"owner"`
	Name      string `gorm:"not null" json:"name"`
	Author    string `gorm:"not null" json:"author"`
	Publisher string `gorm:"not null" json:"publisher"`
	ISBN      string `gorm:"not null" json:"isbn"`
	Price     int    `gorm:"not null" json:"price"`
	Stock     int    `gorm:"not null" json:"stock"`
}
