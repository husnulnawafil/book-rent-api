package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Owner     uint           `gorm:"not null" json:"owner"`
	Name      string         `gorm:"not null" json:"name"`
	Author    string         `gorm:"not null" json:"author"`
	Publisher string         `gorm:"not null" json:"publisher"`
	ISBN      string         `gorm:"not null" json:"isbn"`
	Price     int            `gorm:"not null" json:"price"`
	IsRent    bool           `gorm:"not null" json:"isRent"`
}
