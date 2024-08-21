package handlers

import (
	"fmt"

	"github.com/keliMuthengi/invoiving-api/repo/models"
	"github.com/keliMuthengi/invoiving-api/services"
)

func SendEmail(invoice models.Invoice, tenant models.User, status string) error {

	var subject string
	tenantname := tenant.Username
	invoiceAmount := invoice.InvoiceBalance

	if status != "PAID" {
		subject = fmt.Sprintf("Dear %v your rent bill for this month is %v", tenantname, invoiceAmount)
	}else{
		subject = fmt.Sprintf("Dear %v we have recieved your bill payment of this month of Ksh %v", tenantname, invoiceAmount)
	}
	

	// subject := fmt.Sprintf("Dear %v your rent bill for this month is %v", tenantname, invoiceAmount)
	emailData := services.EmailMessage{
		RecepientEmail: tenant.Email,
		Subject:        subject,
	}

	err := services.SendMail(emailData)

	if err != nil {
		return err
	}
	return nil
}
