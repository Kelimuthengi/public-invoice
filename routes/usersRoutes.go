package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
	"github.com/keliMuthengi/invoiving-api/handlers"
	"github.com/keliMuthengi/invoiving-api/middleware"
)

var (
	auth = middleware.AuthenticationMiddleware()
)

func UserRoutes(g *gin.Engine) {

	userGroup := g.Group("/users")
	{
		userGroup.POST("/addParent", handlers.AddParent)
		userGroup.GET("/listusers",  controllers.DoGetUsers)
		userGroup.POST("/createuser", auth, controllers.CreateUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.GET("/listtenants",controllers.GetTenants)
	}
}
