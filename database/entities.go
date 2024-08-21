package database

import "github.com/keliMuthengi/invoiving-api/repo/models"

var (
	Entities = []interface{}{
		models.Parent{},
		models.Student{},
		models.User{},
		models.Product{},
		models.Invoice{},
		models.InvoiceProduct{},
		models.Transaction{},
		models.HouseUnitName{},
		models.HouseUnitTypes{},
	}
)
