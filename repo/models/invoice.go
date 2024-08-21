package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	InvoiceNumber string    `gorm:"size:255;not null; unique" json:"invoiceNumber"`
	InvoiceDue    time.Time `gorm:"not null" json:"invoiceDue"`
	// Products      []Product `gorm:"foreignkey:InvoiceID" json:"products"`
	Products       []Product `gorm:"many2many:invoice_products" json:"products"`
	StudentID      uint      `gorm:"not null" json:"studentId,omitempty"`
	ParentID       uint      `gorm:"not null" json:"parentId,omitempty"`
	HouseNo        string    `gorm:"size:255" json:"houseno,omitempty"`
	MeterReading   float64   `json:"meterreading,omitempty"`
	Amount         float64   `gorm:"not null" json:"amount"`
	UserId         uint      `gorm:"not null" json:"userid"`
	InvoiceBalance float64   `gorm:"not null" json:"invoicebalance"`
	InvoiceStatus  string    `gorm:"oneof PAID,UNPAID,PARTIAL;not null;default:UNPAID" json:"invoicestatus"`
}
