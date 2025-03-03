package controllers

import (
	"demo/src/orders/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase *application.CreateOrder
}

func NewCreateOrderController(useCase *application.CreateOrder) *CreateOrderController {
	return &CreateOrderController{useCase: useCase}
}

func (co_c *CreateOrderController) Execute(c *gin.Context) {
	var request struct {
		IdProduct int `json:"idProduct" binding:"required"`
		Quantity  int `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := co_c.useCase.Execute(request.IdProduct, request.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Orden creada correctamente"})
}
