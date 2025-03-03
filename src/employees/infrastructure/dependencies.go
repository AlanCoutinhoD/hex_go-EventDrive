package infrastructure

import (
	"demo/src/employees/application"
	"demo/src/employees/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	employeeRepo := NewMySQL()

	getAllEmployee := application.NewListEmployee(employeeRepo)
	listEmployeeController := controllers.NewListEmployeeController(*getAllEmployee)

	createEmployee := application.NewCreateEmployee(employeeRepo)
	createEmployeeController := controllers.NewCreateEmployeeController(*createEmployee)

	updateEmployee := application.NewUpdateEmployee(employeeRepo)
	updateEmployeeController := controllers.NewUpdateEmployeeController(updateEmployee)

	deleteEmployee := application.NewDeleteEmployee(employeeRepo)
	deleteEmployeeController := controllers.NewDeleteEmployeeController(deleteEmployee)

	listEmployeeById := application.NewListEmployeeById(employeeRepo)
	listEmployeeByIdController := controllers.NewListEmployeeByIdController(listEmployeeById)

	employeeRouter := NewEmployeeRouter(listEmployeeController, createEmployeeController, updateEmployeeController, deleteEmployeeController, listEmployeeByIdController)
	employeeRouter.SetupRoutes(router)
}
