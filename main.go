package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keliMuthengi/invoiving-api/env"
	"github.com/keliMuthengi/invoiving-api/routes"
)

func main() {

	// create a new engine instance
	// engine := Engine{}

	trustedproxies := []string{"41.90.43.26"}

	// if err := engine.SetTrustedProxies(trustedproxies); err != nil {
	// 	fmt.Println("Error setting trusted proxies")
	// }
	// list all routes

	// set modes
	env := env.NewEnv()
	mode := env.Mode
	if mode == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}
	

	router := gin.Default()
	if err := router.SetTrustedProxies(trustedproxies); err != nil {
		fmt.Println("Error setting trusted proxies")
		panic("Error setting trusted proxies")
	}

	// apply cors;
	router.Use(cors.Default())

	routes.UserRoutes(router)
	routes.ProductRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TransactionRouter(router)
	routes.UnitsRouter(router)
	routes.Mailroutes(router)

	router.Run()
}
