package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
)

func TransactionRouter(g *gin.Engine) {

	transactionGroup := g.Group("/transaction")
	{
		transactionGroup.POST("/create", controllers.CreateTransaction)
		transactionGroup.GET("/list", controllers.ListTransactions)
	}
}
