package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCase *application.UpdateProduct
}

func NewUpdateProductController(useCase *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCase: useCase}
}

func (up_c *UpdateProductController) Execute(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extrae el ID de la URL
	id := c.Param("id")
	log.Printf("ID recibido de la URL: %s", id)
	product.ID, _ = strconv.Atoi(id) // Convierte el ID de string a int
	log.Printf("ID recibido para actualización: %d", product.ID)

	if err := up_c.useCase.ExecuteByID(
		strconv.Itoa(product.ID),
		product.Name,
		product.Price,
		product.Description,
		product.Image,
	); err != nil {
		log.Printf("Error al actualizar producto: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado correctamente"})
}
