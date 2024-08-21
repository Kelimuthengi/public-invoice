package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID         uint      `gorm:"not null" json:"userid"`
	TransactionRef string    `gorm:"not null;unique" json:"transRef"`
	Products       []Product `gorm:"many2many:transaction_products" json:"products"`
	InvoiceNumber  string    `gorm:"not null" json:"invoiceNumber"`
	Amount         float64   `gorm:"not null" json:"amount"`
	StudentID      uint      `gorm:"not null" json:"studentId"`
	ParentID       uint      `gorm:"not null" json:"parentId"`
	HouseNo        string    `gorm:"size:255" json:"houseno"`
}
