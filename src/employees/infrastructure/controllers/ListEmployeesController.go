package controllers

import (
	"github.com/gin-gonic/gin"
	"demo/src/employees/application"
)

type ListEmployeeController struct {
	useCase application.ListEmployee
}


func NewListEmployeeController(useCase application.ListEmployee) *ListEmployeeController {
	return &ListEmployeeController{useCase: useCase}
}

func (le *ListEmployeeController) Execute(c *gin.Context) {
	employees := le.useCase.Execute()
	c.JSON(200, employees)
}

