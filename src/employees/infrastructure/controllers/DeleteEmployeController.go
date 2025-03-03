package controllers

import (
	"demo/src/employees/application"
	"demo/src/employees/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeController struct {
	useCase *application.DeleteEmployee
}

func NewDeleteEmployeeController(useCase *application.DeleteEmployee) *DeleteEmployeeController {
	return &DeleteEmployeeController{useCase: useCase}
}

func (de *DeleteEmployeeController) Execute(c *gin.Context) {
	id := c.Param("id")

	// Convertir el ID de string a int
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	employee := entities.Employee{ID: intID}

	err = de.useCase.Execute(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo eliminar el empleado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empleado eliminado correctamente"})
}
