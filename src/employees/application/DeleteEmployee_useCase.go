package application

import (
	"demo/src/employees/domain"
	"demo/src/employees/domain/entities"
)

type DeleteEmployee struct {
	employeeRepo domain.IEmployeeRepository
}

func NewDeleteEmployee(repo domain.IEmployeeRepository) *DeleteEmployee {
	return &DeleteEmployee{employeeRepo: repo}
}

func (de *DeleteEmployee) Execute(employee entities.Employee) error {
	return de.employeeRepo.Delete(employee)
}
