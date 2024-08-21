package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/controllers"
)

func UnitsRouter(g *gin.Engine) {

	unitsGroup := g.Group("units")
	{
		unitsGroup.POST("/addunitname", controllers.AddUnitName)
		unitsGroup.POST("/addunit", controllers.AddUnits)
		unitsGroup.GET("/listunits", controllers.Listunitnames)
		unitsGroup.GET("/listhousinguints", controllers.ListHousingunits)
	}
}
