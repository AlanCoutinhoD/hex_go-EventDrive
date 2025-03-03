package application

import (
	"demo/src/employees/domain"
	"log"
	"strconv"
)

type UpdateEmployee struct {
	employeeRepo domain.IEmployeeRepository
}

func NewUpdateEmployee(repo domain.IEmployeeRepository) *UpdateEmployee {
	return &UpdateEmployee{employeeRepo: repo}
}

func (ue *UpdateEmployee) Execute(id string, name string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error al convertir ID: %v", err)
		return err
	}

	log.Printf("Buscando empleado con ID: %d", intID)
	employee, err := ue.employeeRepo.FindByID(intID)
	if err != nil {
		log.Printf("Error al encontrar empleado: %v", err)
		return err
	}

	employeeEntity := *employee
	employeeEntity.Name = name
	log.Printf("Actualizando empleado: %+v", employeeEntity)
	return ue.employeeRepo.Update(employeeEntity)
}
