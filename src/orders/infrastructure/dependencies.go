package infrastructure

import (
	"demo/src/orders/application"
	"demo/src/orders/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	orderRepo := NewMySQL()

	createOrder := application.NewCreateOrder(orderRepo)
	createOrderController := controllers.NewCreateOrderController(createOrder)

	orderRouter := NewOrderRouter(createOrderController)
	orderRouter.SetupRoutes(router)
}
