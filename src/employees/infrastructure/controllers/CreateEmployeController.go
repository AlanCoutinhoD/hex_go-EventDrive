package controllers

import (
	"demo/src/employees/application"
	"demo/src/employees/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeController struct {
	useCase application.CreateEmployee
}

func NewCreateEmployeeController(useCase application.CreateEmployee) *CreateEmployeeController {
	return &CreateEmployeeController{useCase: useCase}
}

func (ce *CreateEmployeeController) Execute(c *gin.Context) {
	var employee entities.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ce.useCase.Execute(employee)
	c.JSON(201, gin.H{"message": "Empleado creado correctamente"})
}
