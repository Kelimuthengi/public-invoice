package handlers

import (
	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func DoCreateTransaction(trans models.Transaction) (models.Transaction, error) {

	if err := database.DB.Create(&trans).Error; err != nil {
		return trans, err
	}

	return trans, nil
}

func DoListTransactions() ([]models.Transaction, error) {

	var transactions []models.Transaction

	if err := database.DB.Model(&models.Transaction{}).Preload("Products").Find(&transactions).Error; err != nil {

		return []models.Transaction{}, nil
	}
	return transactions, nil
}
