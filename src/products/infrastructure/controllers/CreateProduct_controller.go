package controllers

import (
	"demo/src/core"
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase     *application.CreateProduct
	fileService *core.FileService
}

func NewCreateProductController(useCase *application.CreateProduct) *CreateProductController {
	return &CreateProductController{
		useCase:     useCase,
		fileService: core.NewFileService("uploads"),
	}
}

func (cp_c *CreateProductController) Execute(c *gin.Context) {
	// Obtener los datos del formulario
	name := c.PostForm("name")
	priceStr := c.PostForm("price")
	description := c.PostForm("description")

	// Convertir precio a float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Precio inválido"})
		return
	}

	// Manejar la imagen
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar la imagen"})
		return
	}

	// Guardar la imagen
	imagePath, err := cp_c.fileService.SaveFile(file, "products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la imagen"})
		return
	}

	// Crear el producto
	product := entities.NewProduct(name, price, description, imagePath)
	err = cp_c.useCase.Execute(*product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto creado correctamente",
		"product": product,
	})
}
