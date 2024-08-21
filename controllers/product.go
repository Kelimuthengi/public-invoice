package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/repo/models"
)

func AddProduct(c *gin.Context) {
	var productInput handlers.ProductInput

	// Bind JSON input to productInput and handle errors
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{
		ProductName: productInput.ProductName,
		ProductType: productInput.ProductType,
		Price:       productInput.Price,
		Discount:    productInput.Discount,
		UserID:      productInput.UserId,
		Description: productInput.Description,
	}

	data, err := handlers.DoCreateProduct(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"data": data})
}

func ListProducts(c *gin.Context) {

	pagestr := c.Query("page")
	limitstr := c.Query("limit")

	page, err := strconv.Atoi(pagestr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitstr)

	if err != nil || limit < 1 {
		limit = 10
	}

	// get the offset;
	page = (page - 1) * limit

	params := handlers.RequestParams{
		Page:  page,
		Limit: limit,
	}

	products, err := handlers.DoListProducts(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"products": products})

}
