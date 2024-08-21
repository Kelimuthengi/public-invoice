package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

var (
	invoice models.Invoice
)

func CreateInvoice(c *gin.Context) {
	var invoiceinput handlers.InvoiceInput
	var student models.Student
	var parent models.Parent
	var units models.HouseUnitTypes
	var user models.User
	// bind json data;
	if err := c.ShouldBindJSON(&invoiceinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create products
	// get eixsiting student;
	if invoiceinput.StudentID != 0 {
		studentId := uint(invoiceinput.StudentID)
		if err := database.DB.First(&student, studentId).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// get existing parent;
	if invoiceinput.ParentID != 0 {
		parentId := uint(invoiceinput.ParentID)
		if err := database.DB.First(&parent, parentId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// check existing house name;
	if invoiceinput.HouseNo != "" {
		if err := database.DB.Find(&units, "house_no = ?", invoiceinput.HouseNo).Error; err != nil {
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
			return
		}
	}

	// get user for userId::

	if err := database.DB.Find(&user, uint(invoiceinput.UserId)).Error; err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}
	// create invoice products by Id;
	// create an empty array
	var prodslice []models.Product
	var invoiceAmount float64
	for _, invoiceProd := range invoiceinput.Products {
		var product models.Product
		// check product exists;
		if err := database.DB.First(&product, uint(invoiceProd.Id)).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
			return
		}
		invoiceAmount += float64(product.Price)
		prodslice = append(prodslice, product)
	}

	// now pass detsils to DO create invoice function
	invoicedetails := models.Invoice{
		InvoiceNumber:  invoiceinput.InvoiceNumber,
		Products:       prodslice,
		StudentID:      invoiceinput.StudentID,
		ParentID:       invoiceinput.ParentID,
		HouseNo:        invoiceinput.HouseNo,
		MeterReading:   invoiceinput.MeterReading,
		Amount:         invoiceAmount,
		UserId:         user.ID,
		InvoiceBalance: invoiceAmount,
	}

	if invoiceinput.MeterReading != 0 {
		invoicedetails.MeterReading = invoiceinput.MeterReading
	}

	invoiceData, err := handlers.DocreateInvoice(invoicedetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send email;
	err = handlers.SendEmail(invoiceData, user, "UNPAID")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, handlers.ResponseHandler{Data: invoiceData, Message: "Invoice created and message sent to email successfully!"})
}

func GetInvoices(c *gin.Context) {

	// get all pagination related stuff here!

	allinvoices, err := handlers.DoGetInvoices()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"data": allinvoices})
}

func GetInvoiceStatus(invoiceNumber string, Amount float64) error {

	if err := database.DB.Find(&invoice, "invoice_number = ?", invoiceNumber).Error; err != nil {
		return err
	}

	invoicebalance := invoice.InvoiceBalance
	// checks Invoice status
	if invoicebalance == 0 {
		return errors.New("invoice has been fully settled")
	}

	if Amount > invoicebalance {
		return errors.New("invoice overpayment is not supported")
	}
	return nil
}

func SettleInvoice(invoiceNumber string, Amount float64) error {

	invoicebalance := invoice.InvoiceBalance
	if err := database.DB.First(&invoice, "invoice_number = ?", invoiceNumber).Error; err != nil {
		return err
	}
	invoicebalance = (invoicebalance - Amount)

	if invoicebalance == 0 {
		invoice.InvoiceStatus = "PAID"
		invoice.InvoiceBalance = 0
	} else {
		invoice.InvoiceStatus = "PARTIAL"
		invoice.InvoiceBalance = invoicebalance
	}

	if err := database.DB.Save(&invoice).Error; err != nil {
		return err
	}
	return nil
}
