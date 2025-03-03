package infrastructure

import (
	"demo/src/orders/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
	createOrderController *controllers.CreateOrderController
}

func NewOrderRouter(createOrderController *controllers.CreateOrderController) *OrderRouter {
	return &OrderRouter{
		createOrderController: createOrderController,
	}
}

func (or *OrderRouter) SetupRoutes(router *gin.Engine) {
	ordersGroup := router.Group("/orders")
	{
		ordersGroup.POST("", or.createOrderController.Execute)
	}
}
