package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/repo/models"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	var transactionRequest handlers.TransactionRequest
	var student models.Student
	var user models.User
	var invoice models.Invoice
	var studentId uint
	// bind json data:;;
	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Status: 1, Message: err.Error()})
		return
	}
	// build transaction body here; check for existence of products;
	var productSlice = []models.Product{}
	for _, inproduct := range transactionRequest.Products {
		var product models.Product

		if err := database.DB.First(&product, uint(inproduct.ProductId)).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				notfoundmessage := fmt.Sprintf("No product with ID %v", strconv.FormatUint(uint64(inproduct.ProductId), 10))
				c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: notfoundmessage, Status: 1})
				return
			}
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
			return
		}

		productSlice = append(productSlice, product)
	}

	fmt.Println(transactionRequest.InvoiceNumber)
	// get student details from invoiceNumber;
	if err := database.DB.Find(&invoice, "invoice_number = ?", transactionRequest.InvoiceNumber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// notfoundmessage := fmt.Sprintf("No student with ID %v", strconv.FormatUint(uint64(studentId), 10))
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
			return
		}
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}
	studentId = invoice.StudentID

	if studentId != 0 {
		// get parent details here;
		if err := database.DB.Find(&student, studentId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				notfoundmessage := fmt.Sprintf("No student with ID %v", strconv.FormatUint(uint64(studentId), 10))
				c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: notfoundmessage, Status: 1})
				return
			}
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
			return

		}

		// get user details from parentID
		if err := database.DB.Find(&user, student.ParentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				notfoundmessage := fmt.Sprintf("No user with ID %v", strconv.FormatUint(uint64(student.ParentID), 10))
				c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: notfoundmessage, Status: 1})
				return
			}
			c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
			return
		}
	}

	// get user details form invoice;
	if err := database.DB.Find(&user, invoice.UserId).Error; err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	transaction := models.Transaction{
		Amount:         transactionRequest.Amount,
		StudentID:      student.ID,
		ParentID:       student.ParentID,
		UserID:         user.ID,
		Products:       productSlice,
		InvoiceNumber:  transactionRequest.InvoiceNumber,
		TransactionRef: transactionRequest.TransactionRef,
		HouseNo:        invoice.HouseNo,
	}

	// invoice checks before transactions;
	err := GetInvoiceStatus(transactionRequest.InvoiceNumber, transactionRequest.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	trans, err := handlers.DoCreateTransaction(transaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	err = SettleInvoice(invoice.InvoiceNumber, transactionRequest.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	// send invoice to Tenant
	err = handlers.SendEmail(invoice, user, "PAID")

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}
	c.JSON(http.StatusOK, handlers.ResponseHandler{Message: "Transaction created successfully and email sent to customer", Data: trans})

}

func ListTransactions(c *gin.Context) {

	// get transactions;

	transactions, err := handlers.DoListTransactions()

	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.ResponseHandler{Message: err.Error(), Status: 1})
		return
	}

	c.JSON(http.StatusOK, handlers.ResponseHandler{Data: transactions, Message: "Data retrieved successfully", Status: 0})
}
