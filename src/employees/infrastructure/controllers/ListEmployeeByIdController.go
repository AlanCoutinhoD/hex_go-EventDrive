package controllers

import (
	"demo/src/employees/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListEmployeeByIdController struct {
	useCase *application.ListEmployeeById
}

func NewListEmployeeByIdController(useCase *application.ListEmployeeById) *ListEmployeeByIdController {
	return &ListEmployeeByIdController{useCase: useCase}
}

func (c *ListEmployeeByIdController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")
	employee, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, employee)
}
