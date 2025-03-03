package controllers

import (
	"demo/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCase *application.DeleteProduct
}

func NewDeleteProductController(useCase *application.DeleteProduct) *DeleteProductController {
	return &DeleteProductController{useCase: useCase}
}

func (dp_c *DeleteProductController) Execute(c *gin.Context) {
	// Extrae el ID de la URL
	id := c.Param("id")
	intID, err := strconv.Atoi(id) // Convierte el ID de string a int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llama al caso de uso para eliminar el producto usando el ID
	if err := dp_c.useCase.Execute(intID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
