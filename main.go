package main

import (
	employeesInfrastructure "demo/src/employees/infrastructure"
	ordersInfrastructure "demo/src/orders/infrastructure"
	productsInfrastructure "demo/src/products/infrastructure"

	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	router := gin.Default()

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://98.83.96.54:3000"}
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

	serverAddr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("🚀 Servidor corriendo en http://%s", serverAddr)
	log.Printf("📝 Endpoints disponibles:")
	log.Printf("   POST http://%s/orders", serverAddr)
	log.Printf("   GET  http://%s/orders", serverAddr)

	router.Run(":" + port) // Importante: usar :port para escuchar en todas las interfaces
}
