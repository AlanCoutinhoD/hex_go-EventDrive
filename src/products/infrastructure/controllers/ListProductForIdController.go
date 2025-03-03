package controllers

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListProductForIdController struct {
	useCase *application.ListProductForId
}

func NewListProductForIdController(useCase *application.ListProductForId) *ListProductForIdController {
	return &ListProductForIdController{useCase: useCase}
}

func (c *ListProductForIdController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	// Construir la URL base
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := scheme + "://" + ctx.Request.Host

	// Crear la respuesta con la URL completa de la imagen
	response := ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		ImageURL:    baseURL + "/uploads" + product.Image,
	}

	ctx.JSON(http.StatusOK, response)
}
