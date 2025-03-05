package infrastructure

import (
	"demo/src/orders/application"
	"demo/src/orders/infrastructure/controllers"
	"demo/src/orders/infrastructure/rabbitmq"

	"log"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	orderRepo := NewMySQL()
	messageBroker, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}

	createOrder := application.NewCreateOrder(orderRepo, messageBroker)
	createOrderController := controllers.NewCreateOrderController(createOrder)

	orderRouter := NewOrderRouter(createOrderController)
	orderRouter.SetupRoutes(router)
}
