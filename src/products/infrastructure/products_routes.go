package infrastructure

import (
	"demo/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	listController        *controllers.ListProductController
	createController      *controllers.CreateProductController
	listForIdController   *controllers.ListProductForIdController
	deleteController      *controllers.DeleteProductController
	updateController      *controllers.UpdateProductController
	uploadImageController *controllers.UploadProductImageController
}

func NewProductRouter(listController *controllers.ListProductController, createController *controllers.CreateProductController, listForIdController *controllers.ListProductForIdController, deleteController *controllers.DeleteProductController, updateController *controllers.UpdateProductController, uploadImageController *controllers.UploadProductImageController) *ProductRouter {
	return &ProductRouter{
		listController:        listController,
		createController:      createController,
		listForIdController:   listForIdController,
		deleteController:      deleteController,
		updateController:      updateController,
		uploadImageController: uploadImageController,
	}
}

func (pr *ProductRouter) SetupRoutes(router *gin.Engine) {
	productsGroup := router.Group("/products")
	{
		productsGroup.GET("", pr.listController.Execute)
		productsGroup.GET("/:id", pr.listForIdController.Execute)
		productsGroup.POST("", pr.createController.Execute)
		productsGroup.DELETE("/:id", pr.deleteController.Execute)
		productsGroup.PUT("/:id", pr.updateController.Execute)
		productsGroup.POST("/upload-image", pr.uploadImageController.Execute)
	}

	// Servir archivos estáticos
	router.Static("/uploads", "./uploads")
}
