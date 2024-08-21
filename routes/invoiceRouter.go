package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
)

func InvoiceRoutes(c *gin.Engine) {

	invoiceGroup := c.Group("/invoices")
	{
		invoiceGroup.POST("/createinvoice", controllers.CreateInvoice)
		invoiceGroup.GET("/listinvoices", controllers.GetInvoices)
	}
}
