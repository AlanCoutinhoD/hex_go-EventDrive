package infrastructure

import (
	"demo/src/employees/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type EmployeeRouter struct {
	listController         *controllers.ListEmployeeController
	createController       *controllers.CreateEmployeeController
	updateController       *controllers.UpdateEmployeeController
	deleteController       *controllers.DeleteEmployeeController
	listByIdController     *controllers.ListEmployeeByIdController
}

func NewEmployeeRouter(
	listController *controllers.ListEmployeeController,
	createController *controllers.CreateEmployeeController,
	updateController *controllers.UpdateEmployeeController,
	deleteController *controllers.DeleteEmployeeController,
	listByIdController *controllers.ListEmployeeByIdController,
) *EmployeeRouter {
	return &EmployeeRouter{
		listController:     listController,
		createController:   createController,
		updateController:   updateController,
		deleteController:   deleteController,
		listByIdController: listByIdController,
	}
}

func (er *EmployeeRouter) SetupRoutes(router *gin.Engine) {
	employeesGroup := router.Group("/employees")
	{
		employeesGroup.POST("", er.createController.Execute)
		employeesGroup.PUT("/:id", er.updateController.Execute)
		employeesGroup.DELETE("/:id", er.deleteController.Execute)
		employeesGroup.GET("", er.listController.Execute)
		employeesGroup.GET("/:id", er.listByIdController.Execute)
	}
}


