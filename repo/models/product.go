package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string    `gorm:"size:255;not null; unique" json:"productName"`
	ProductType string    `gorm:"size:255;not null; default:service" json:"productType"`
	Description string    `gorm:"size:400; not null" json:"description"`
	Price       int       `gorm:"not null;" json:"price"`
	Discount    int       `gorm:"not null; default:0" json:"discount"`
	UserID      uint      `gorm:"not null" json:"userid"`
}
