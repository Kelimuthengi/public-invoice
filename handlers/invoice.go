package handlers

import (
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func DocreateInvoice(inv models.Invoice) (models.Invoice, error) {

	if err :=  database.DB.Create(&inv).Error; err != nil {
		return inv, err
	}

	return inv, nil
}

func DoGetInvoices() ([]models.Invoice, error) {

	var invoices []models.Invoice

	if err := database.DB.Model(&models.Invoice{}).Preload("Products").Find(&invoices).Error; err != nil {
		return invoices, err
	}
	return invoices, nil
}
