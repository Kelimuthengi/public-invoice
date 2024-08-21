package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
)

func ProductRoutes(g *gin.Engine) {
	productGroup := g.Group("/product")
	{
		productGroup.POST("/add_product", controllers.AddProduct)
		productGroup.GET("/listproducts", controllers.ListProducts)
	}

}
