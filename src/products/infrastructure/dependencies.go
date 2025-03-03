package infrastructure

import (
	"demo/src/products/application"
	"demo/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Crear producto
	ps := NewMySQL()
	createProduct := application.NewCreateProduct(ps)
	createProductController := controllers.NewCreateProductController(createProduct)

	// Listar todos los productos
	productRepo := ps
	getAllProduct := application.NewListProduct(productRepo)
	listProductController := controllers.NewListProductController(*getAllProduct)

	// Listar producto por id
	listProductForId := application.NewListProductForId(productRepo)
	listProductForIdController := controllers.NewListProductForIdController(listProductForId)

	// Eliminar producto
	deleteProduct := application.NewDeleteProduct(productRepo)
	deleteProductController := controllers.NewDeleteProductController(deleteProduct)

	// Actualizar producto
	updateProduct := application.NewUpdateProduct(productRepo)
	updateProductController := controllers.NewUpdateProductController(updateProduct)

	uploadImageController := controllers.NewUploadProductImageController()

	productRouter := NewProductRouter(listProductController, createProductController, listProductForIdController, deleteProductController, updateProductController, uploadImageController)
	productRouter.SetupRoutes(router)
}
