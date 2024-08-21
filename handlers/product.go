package handlers

import (
	"errors"

	"github.com/keliMuthengi/invoiving-api/database"
	"github.com/keliMuthengi/invoiving-api/repo/models"
	"gorm.io/gorm"
)

func DoCreateProduct(p models.Product) (models.Product, error) {

	product := models.Product{
		ProductName: p.ProductName,
		ProductType: p.ProductType,
		Description: p.Description,
		Price:       p.Price,
		Discount:    p.Discount,
		UserID:      p.UserID,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return product, errors.New("duplicate keys")
		}
		return product, errors.New("error creating product")
	}
	return product, nil
}

func DoListProducts(params RequestParams) ([]models.Product, error) {

	var products []models.Product

	if err := database.DB.Find(&products).Preload("Users").Error; err != nil {
		return products, err
	}
	return products, nil
}
