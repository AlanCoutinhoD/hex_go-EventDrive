package controllers

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListProductController struct {
	useCase application.ListProduct
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func NewListProductController(useCase application.ListProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (lp_c *ListProductController) Execute(c *gin.Context) {
	products := lp_c.useCase.Execute()

	// Construir la URL base
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := scheme + "://" + c.Request.Host

	// Transformar los productos a la respuesta con URLs completas
	var response []ProductResponse
	for _, product := range products {
		response = append(response, ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			ImageURL:    baseURL + "/uploads" + product.Image,
		})
	}

	c.JSON(http.StatusOK, response)
}
