package controllers

import (
	"demo/src/employees/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeController struct {
	updateEmployeeUseCase *application.UpdateEmployee
}

func NewUpdateEmployeeController(useCase *application.UpdateEmployee) *UpdateEmployeeController {
	return &UpdateEmployeeController{updateEmployeeUseCase: useCase}
}

func (c *UpdateEmployeeController) Execute(ctx *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	id := ctx.Param("id")
	log.Printf("ID recibido para actualización: %s", id)
	if err := c.updateEmployeeUseCase.Execute(id, input.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update employee"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}
