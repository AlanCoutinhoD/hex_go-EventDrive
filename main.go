package main

import (
	employeesInfrastructure "demo/src/employees/infrastructure"
	ordersInfrastructure "demo/src/orders/infrastructure"
	productsInfrastructure "demo/src/products/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	// Inicializar infraestructura de productos
	productsInfrastructure.Init(router)

	// Inicializar infraestructura de empleados
	employeesInfrastructure.Init(router)

	// Inicializar infraestructura de órdenes
	ordersInfrastructure.Init(router)

	router.Run(":8080")
}
